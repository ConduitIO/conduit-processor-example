// Copyright © 2024 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package complexp

import (
	"context"
	"fmt"
	"strconv"

	"github.com/conduitio/conduit-commons/config"
	"github.com/conduitio/conduit-commons/opencdc"
	sdk "github.com/conduitio/conduit-processor-sdk"
)

//go:generate paramgen -output=processorConfig_paramgen.go processorConfig

func NewProcessor() sdk.Processor {
	return &Processor{}
}

type Processor struct {
	referenceResolver sdk.ReferenceResolver
	config            processorConfig

	sdk.UnimplementedProcessor
}

type processorConfig struct {
	// Field is the target field that will be set.
	Field string `json:"field" validate:"required,exclusion=.Position"`
	// Threshold is the threshold for filtering the record.
	Threshold int `json:"threshold" validate:"required,gt=0"`
}

func (p *Processor) Specification() (sdk.Specification, error) {
	return sdk.Specification{
		Name:       "processor-complex",
		Version:    "v1.0.0",
		Parameters: processorConfig{}.Parameters(),
	}, nil
}

func (p *Processor) Configure(ctx context.Context, c config.Config) error {
	err := sdk.ParseConfig(ctx, c, &p.config, processorConfig{}.Parameters())
	if err != nil {
		return fmt.Errorf("failed to parse configuration: %w", err)
	}

	resolver, err := sdk.NewReferenceResolver(p.config.Field)
	if err != nil {
		return fmt.Errorf("failed to parse the %q param: %w", "field", err)
	}
	p.referenceResolver = resolver
	return nil
}

func (p *Processor) Process(ctx context.Context, records []opencdc.Record) []sdk.ProcessedRecord {
	output := make([]sdk.ProcessedRecord, 0, len(records))
	for _, record := range records {
		rec := record
		ref, err := p.referenceResolver.Resolve(&rec)
		if err != nil {
			return append(output, sdk.ErrorRecord{Error: err})
		}

		var rawStr string
		switch raw := ref.Get().(type) {
		case opencdc.RawData:
			rawStr = string(raw)
		default:
			rawStr = fmt.Sprintf("%v", raw)
		}

		sdk.Logger(ctx).Info().Msgf("comparing %v to threshold (%v)", rawStr, p.config.Threshold)

		var thresholdEvaluation string
		if value, err := strconv.Atoi(rawStr); err != nil {
			sdk.Logger(ctx).Warn().Msgf("ignoring record, %q is not an integer", rawStr)
			thresholdEvaluation = "n/a"
		} else if value > p.config.Threshold {
			thresholdEvaluation = "true"
		} else {
			thresholdEvaluation = "false"
		}
		rec.Metadata["over_threshold"] = thresholdEvaluation

		output = append(output, sdk.SingleRecord(rec))
	}
	return output
}

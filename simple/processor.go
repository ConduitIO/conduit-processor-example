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

package simple

import (
	"context"

	"github.com/conduitio/conduit-commons/opencdc"
	sdk "github.com/conduitio/conduit-processor-sdk"
)

func NewProcessor() sdk.Processor {
	return sdk.NewProcessorFunc(spec, Process)
}

var spec = sdk.Specification{
	Name:    "processor-simple",
	Version: "v1.0.0",
}

func Process(ctx context.Context, record opencdc.Record) (opencdc.Record, error) {
	if record.Metadata["over_threshold"] == "true" {
		sdk.Logger(ctx).Warn().Msg("filtering record because it is over the threshold")
		return opencdc.Record{}, sdk.ErrFilterRecord
	}
	record.Metadata["processed-by"] = "processor-simple"
	return record, nil
}

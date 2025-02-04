// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"context"
	"fmt"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type SubstringArguments[K any] struct {
	Target ottl.StringGetter[K] `ottlarg:"0"`
	Start  int64                `ottlarg:"1"`
	Length int64                `ottlarg:"2"`
}

func NewSubstringFactory[K any]() ottl.Factory[K] {
	return ottl.NewFactory("Substring", &SubstringArguments[K]{}, createSubstringFunction[K])
}

func createSubstringFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	args, ok := oArgs.(*SubstringArguments[K])

	if !ok {
		return nil, fmt.Errorf("SubstringFactory args must be of type *SubstringArguments[K]")
	}

	return substring(args.Target, args.Start, args.Length)
}

func substring[K any](target ottl.StringGetter[K], start int64, length int64) (ottl.ExprFunc[K], error) {
	if start < 0 {
		return nil, fmt.Errorf("invalid start for substring function, %d cannot be negative", start)
	}
	if length <= 0 {
		return nil, fmt.Errorf("invalid length for substring function, %d cannot be negative or zero", length)
	}

	return func(ctx context.Context, tCtx K) (interface{}, error) {
		val, err := target.Get(ctx, tCtx)
		if err != nil {
			return nil, err
		}
		if (start + length) > int64(len(val)) {
			return nil, fmt.Errorf("invalid range for substring function, %d cannot be greater than the length of target string(%d)", start+length, len(val))
		}
		return val[start : start+length], nil
	}, nil
}

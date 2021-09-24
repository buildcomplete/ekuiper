// Copyright 2021 EMQ Technologies Co., Ltd.
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

package store

import (
	"fmt"
	"github.com/lf-edge/ekuiper/internal/pkg/db"
	"github.com/lf-edge/ekuiper/internal/pkg/db/redis"
	"github.com/lf-edge/ekuiper/internal/pkg/db/sql"
	rb "github.com/lf-edge/ekuiper/internal/pkg/store/redis"
	sb "github.com/lf-edge/ekuiper/internal/pkg/store/sql"
	"github.com/lf-edge/ekuiper/pkg/kv"
)

type Builder interface {
	CreateStore(table string) (kv.KeyValue, error)
}

func CreateStoreBuilder(database db.Database) (Builder, error) {
	switch database.(type) {
	case *redis.Instance:
		d := *database.(*redis.Instance)
		return rb.NewStoreBuilder(d), nil
	case sql.Database:
		d := database.(sql.Database)
		return sb.NewStoreBuilder(d), nil
	default:
		return nil, fmt.Errorf("unrecognized database type")
	}
}
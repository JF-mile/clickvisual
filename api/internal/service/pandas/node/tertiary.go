package node

import (
	"github.com/pkg/errors"

	"github.com/clickvisual/clickvisual/api/internal/pkg/constx"
	"github.com/clickvisual/clickvisual/api/internal/pkg/model/db"
	"github.com/clickvisual/clickvisual/api/internal/pkg/model/view"
)

type tertiary struct {
	next department
}

func (r *tertiary) execute(n *node) (res view.RunNodeResult, err error) {
	if n.tertiaryDone {
		return
	}
	n.tertiaryDone = true
	switch n.n.Tertiary {
	case db.TertiaryClickHouse:
		return doTyClickHouse(n)
	case db.TertiaryMySQL:
		return doTyMySQL(n)
	case db.TertiaryOfflineSync:
		return doTyOfflineSync(n)
	case db.TertiaryRealTimeSync:
		return doTyRealTimeSync(n)
	default:
		return res, errors.Wrap(constx.ErrBigdataNotSupportNodeType, "tertiary execute")
	}
}

func (r *tertiary) setNext(next department) {
	r.next = next
}

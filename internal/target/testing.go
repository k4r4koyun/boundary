package target

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/boundary/internal/db"
	"github.com/hashicorp/go-uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
)

func TestTcpTarget(t *testing.T, conn *gorm.DB, scopeId, name string, opt ...Option) *TcpTarget {
	t.Helper()
	require := require.New(t)
	rw := db.New(conn)
	target, err := NewTcpTarget(scopeId, name, opt...)
	require.NoError(err)
	id, err := newTcpId()
	require.NoError(err)
	target.PublicId = id
	err = rw.Create(context.Background(), target)
	require.NoError(err)
	return target
}

func testTargetName(t *testing.T, scopeId string) string {
	t.Helper()
	id, err := uuid.GenerateUUID()
	require.NoError(t, err)
	return fmt.Sprintf("%s-%s", scopeId, id)
}
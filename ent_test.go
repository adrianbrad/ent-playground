package ent_playground

import (
	"context"
	"ent-playground/ent"
	"fmt"
	"github.com/adrianbrad/dbutils"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQuery(t *testing.T) {
	require := require.New(t)
	client, err := ent.Open("postgres", dbutils.DataSource{
		DBName:   "ent",
	}.String())
	require.NoError(err)
	ctx := context.Background()
	user, err := client.User.Get(ctx, "as")
	require.NoError(err)
	fmt.Println(user.Name)

	user, err = client.User.Create().SetID("1234").SetEmail("brad@brad").SetName("brad").Save(ctx)
	require.NoError(err)
	fmt.Println(user)
}
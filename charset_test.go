package c4_test

import (
	"testing"

	"github.com/Avalanche-io/c4"
	"github.com/cheekybits/is"
)

func TestCheckCharacterSet(t *testing.T) {
	is := is.New(t)
	tests := []struct {
		A       string
		B       string
		IsError bool
		ID      string
	}{
		{
			A:       "c41111uoSFHtMmbcTLJmFYvnjuA8EAMQgHQbE3zwmzuoM2epGzDeTvjPYeeuHvzGheFqA1x4qe84x7UVPrRxMLUiRr",
			B:       "c41111VPsgiUnMBCtmjMgyWNKVb8fbnqGiqBf3aXMaVPn2EQhaeEtWKpyEEViWahHEgRb1Y4RE84Y7uvpSrYnmuJrS",
			IsError: false,
			ID:      "c41111VPsgiUnMBCtmjMgyWNKVb8fbnqGiqBf3aXMaVPn2EQhaeEtWKpyEEViWahHEgRb1Y4RE84Y7uvpSrYnmuJrS",
		},
		{
			A:       "c41111VPsgiUnMBCtmjMgyWNKVb8fbnqGiqBf3aXMaVPn2EQhaeEtWKpyEEViWahHEgRb1Y4RE84Y7uvpSrYnmuJrS",
			B:       "c41111uoSFHtMmbcTLJmFYvnjuA8EAMQgHQbE3zwmzuoM2epGzDeTvjPYeeuHvzGheFqA1x4qe84x7UVPrRxMLUiRr",
			IsError: false,
			ID:      "c41111VPsgiUnMBCtmjMgyWNKVb8fbnqGiqBf3aXMaVPn2EQhaeEtWKpyEEViWahHEgRb1Y4RE84Y7uvpSrYnmuJrS",
		},
		{
			A:       "c41111uoSFHtMmbcTLJmFYvnjuA8EAMQgHQbE3zwmzuoM2epGzDeTvjPYeeuHvzGheFqA1x4qe84x7UVPrRxMLUiRr",
			B:       "c42fcSBQEaPKmsmwJqr2GFGQbDjMfdhtbxq9WWJCe5aZ2XnQwETF5nkjR3zt5KqcWy88ay6de1NeCXGHP5tgxA4W2t",
			IsError: true,
			ID:      "",
		},
	}

	for _, test := range tests {
		a, _ := c4.ParseID(test.A)
		b, _ := c4.ParseID(test.B)
		id, err := c4.CheckCharacterSet(a, b)
		if test.IsError {
			is.Err(err)
			is.Nil(id)
			continue
		}
		test_id, _ := c4.ParseID(test.ID)
		is.Equal(id.String(), test_id.String())
	}
}

func TestOldCharsetIDToNew(t *testing.T) {
	is := is.New(t)
	newid, _ := c4.ParseID("c41111VPsgiUnMBCtmjMgyWNKVb8fbnqGiqBf3aXMaVPn2EQhaeEtWKpyEEViWahHEgRb1Y4RE84Y7uvpSrYnmuJrS")
	oldid, _ := c4.ParseID("c41111uoSFHtMmbcTLJmFYvnjuA8EAMQgHQbE3zwmzuoM2epGzDeTvjPYeeuHvzGheFqA1x4qe84x7UVPrRxMLUiRr")
	id := c4.OldCharsetIDToNew(oldid)
	is.NotNil(id)
	is.Equal(newid.String(), id.String())
}

func TestNewCharsetIDToOld(t *testing.T) {
	is := is.New(t)
	newid, _ := c4.ParseID("c41111VPsgiUnMBCtmjMgyWNKVb8fbnqGiqBf3aXMaVPn2EQhaeEtWKpyEEViWahHEgRb1Y4RE84Y7uvpSrYnmuJrS")
	oldid, _ := c4.ParseID("c41111uoSFHtMmbcTLJmFYvnjuA8EAMQgHQbE3zwmzuoM2epGzDeTvjPYeeuHvzGheFqA1x4qe84x7UVPrRxMLUiRr")
	id := c4.NewCharsetIDToOld(newid)
	is.NotNil(id)
	is.Equal(oldid.String(), id.String())
}

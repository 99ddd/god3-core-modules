package language

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
)

func TestAdd(t *testing.T) {
	Add("en", map[string]string{})
}

func TestGetWithScope(t *testing.T) {
	config.Set(config.Config{
		Language: EN,
	})
	en["foo"] = "bar"
	assert.Equal(t, GetWithScope("foo"), "bar")
	en["user.table.foo2"] = "bar"
	assert.Equal(t, GetWithScope("foo2"), "foo2")
	assert.Equal(t, GetWithScope("foo2", "user"), "foo2")
	assert.Equal(t, GetWithScope("foo2", "user", "table"), "bar")
}

func TestGet(t *testing.T) {
	config.Set(config.Config{
		Language: EN,
	})
	en["foo"] = "bar"
	assert.Equal(t, Get("foo"), "bar")
}

func TestWithScopes(t *testing.T) {
	assert.Equal(t, WithScopes("foo", "user", "table"), "user.table.foo")
}

func TestGetFromHtml(t *testing.T) {
	config.Set(config.Config{
		Language: EN,
	})
	en["user.table.foo"] = "bar"
	assert.Equal(t, GetFromHtml("foo", "user", "table"), template.HTML("bar"))
}
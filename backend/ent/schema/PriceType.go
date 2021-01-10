package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// PriceType holds the schema definition for the PriceType entity.
type PriceType struct {
	ent.Schema
}

// Fields of the Playlist.
func (PriceType) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
	}
}

// Edges of the PriceType.
func (PriceType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("playlists").Unique(),
		edge.To("playlist_videos", Playlist_Video.Type).StorageKey(edge.Column("playlist_id")),
	}
}

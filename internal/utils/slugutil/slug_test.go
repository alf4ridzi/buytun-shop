package slugutil

import "testing"

func TestGenerateSlugProduct(t *testing.T) {
	text := GenerateProductSlug("Lenovo V14 G4 Ryzen 5 7520U 8GB / 16GB 512GB W11+OHS 14 FHD")
	t.Log(text)
}

// Code generated by protoc-gen-goext. DO NOT EDIT.

package image_generation

func (m *Message) SetText(v string) {
	m.Text = v
}

func (m *Message) SetWeight(v float64) {
	m.Weight = v
}

func (m *AspectRatio) SetWidthRatio(v int64) {
	m.WidthRatio = v
}

func (m *AspectRatio) SetHeightRatio(v int64) {
	m.HeightRatio = v
}

func (m *ImageGenerationOptions) SetMimeType(v string) {
	m.MimeType = v
}

func (m *ImageGenerationOptions) SetSeed(v int64) {
	m.Seed = v
}

func (m *ImageGenerationOptions) SetAspectRatio(v *AspectRatio) {
	m.AspectRatio = v
}

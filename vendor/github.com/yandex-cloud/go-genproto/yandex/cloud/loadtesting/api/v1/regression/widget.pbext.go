// Code generated by protoc-gen-goext. DO NOT EDIT.

package regression

import (
	report "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/report"
)

type Widget_Widget = isWidget_Widget

func (m *Widget) SetWidget(v Widget_Widget) {
	m.Widget = v
}

func (m *Widget) SetPosition(v *Widget_LayoutPosition) {
	m.Position = v
}

func (m *Widget) SetChart(v *ChartWidget) {
	m.Widget = &Widget_Chart{
		Chart: v,
	}
}

func (m *Widget) SetText(v *TextWidget) {
	m.Widget = &Widget_Text{
		Text: v,
	}
}

func (m *Widget) SetTitle(v *TitleWidget) {
	m.Widget = &Widget_Title{
		Title: v,
	}
}

func (m *Widget_LayoutPosition) SetX(v int64) {
	m.X = v
}

func (m *Widget_LayoutPosition) SetY(v int64) {
	m.Y = v
}

func (m *Widget_LayoutPosition) SetWidth(v int64) {
	m.Width = v
}

func (m *Widget_LayoutPosition) SetHeight(v int64) {
	m.Height = v
}

func (m *ChartWidget) SetId(v string) {
	m.Id = v
}

func (m *ChartWidget) SetName(v string) {
	m.Name = v
}

func (m *ChartWidget) SetDescription(v string) {
	m.Description = v
}

func (m *ChartWidget) SetFilterStr(v string) {
	m.FilterStr = v
}

func (m *ChartWidget) SetTestCase(v string) {
	m.TestCase = v
}

func (m *ChartWidget) SetKpis(v []*report.Kpi) {
	m.Kpis = v
}

func (m *TextWidget) SetText(v string) {
	m.Text = v
}

func (m *TitleWidget) SetText(v string) {
	m.Text = v
}

func (m *TitleWidget) SetSize(v TitleWidget_TitleSize) {
	m.Size = v
}

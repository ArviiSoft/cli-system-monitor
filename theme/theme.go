package theme

import (
    "encoding/json"
    "os"
    ui "github.com/gizak/termui/v3"
)

type ThemeConfig struct {
    Theme struct {
        CPUColor  string `json:"cpu_color"`
        RAMColor  string `json:"ram_color"`
        DiskColor string `json:"disk_color"`
    } `json:"theme"`
}

func LoadTheme(path string) (ThemeConfig, error) {
    var config ThemeConfig
    data, err := os.ReadFile(path)
    if err != nil {
        return config, err
    }
    err = json.Unmarshal(data, &config)
    return config, err
}

func GetColor(name string) ui.Color {
    switch name {
    case "red":
        return ui.ColorRed
    case "green":
        return ui.ColorGreen
    case "blue":
        return ui.ColorBlue
    case "cyan":
        return ui.ColorCyan
    case "magenta":
        return ui.ColorMagenta
    case "yellow":
        return ui.ColorYellow
    case "white":
        return ui.ColorWhite
    default:
        return ui.ColorClear
    }
}
package main

import (
    "fmt"
    "log"
    "math/rand"
    "os"
    "time"

    ui "github.com/gizak/termui/v3"
    "github.com/gizak/termui/v3/widgets"

    "cli-system-monitor/theme"
)

func main() {
    if err := ui.Init(); err != nil {
        log.Fatalf("failed to initialize Termui: %v", err)
    }
    defer ui.Close()

    config, err := theme.LoadTheme("config/config.json")
    if err != nil {
        fmt.Println("Error loading config:", err)
        os.Exit(1)
    }

    cpuGauge := widgets.NewGauge()
    cpuGauge.Title = "CPU Usage"
    cpuGauge.Percent = 0
    cpuGauge.BarColor = theme.GetColor(config.Theme.CPUColor)
    cpuGauge.BorderStyle.Fg = ui.ColorWhite
    cpuGauge.TitleStyle.Fg = ui.ColorWhite

    ramGauge := widgets.NewGauge()
    ramGauge.Title = "RAM Usage"
    ramGauge.Percent = 0
    ramGauge.BarColor = theme.GetColor(config.Theme.RAMColor)
    ramGauge.BorderStyle.Fg = ui.ColorWhite
    ramGauge.TitleStyle.Fg = ui.ColorWhite

    diskGauge := widgets.NewGauge()
    diskGauge.Title = "Disk Usage"
    diskGauge.Percent = 0
    diskGauge.BarColor = theme.GetColor(config.Theme.DiskColor)
    diskGauge.BorderStyle.Fg = ui.ColorWhite
    diskGauge.TitleStyle.Fg = ui.ColorWhite

    spark := widgets.NewSparkline()
    spark.LineColor = ui.ColorGreen
    spark.Data = []float64{}

    sparkGroup := widgets.NewSparklineGroup(spark)
    sparkGroup.Title = "CPU History"

    grid := ui.NewGrid()
    termWidth, termHeight := ui.TerminalDimensions()
    grid.SetRect(0, 0, termWidth, termHeight)
    grid.Set(
        ui.NewRow(0.3,
            ui.NewCol(1.0, cpuGauge),
        ),
        ui.NewRow(0.3,
            ui.NewCol(1.0, ramGauge),
        ),
        ui.NewRow(0.3,
            ui.NewCol(1.0, diskGauge),
        ),
        ui.NewRow(0.1,
            ui.NewCol(1.0, sparkGroup),
        ),
    )

    ui.Render(grid)

    uiEvents := ui.PollEvents()
    ticker := time.NewTicker(time.Second).C

    for {
        select {
        case e := <-uiEvents:
            if e.Type == ui.KeyboardEvent {
                return
            }
        case <-ticker:
            cpu := rand.Intn(100)
            ram := rand.Intn(100)
            disk := rand.Intn(100)

            cpuGauge.Percent = cpu
            ramGauge.Percent = ram
            diskGauge.Percent = disk

            spark.Data = append(spark.Data, float64(cpu))
            if len(spark.Data) > 30 {
                spark.Data = spark.Data[1:]
            }

            ui.Render(grid)
        }
    }
}
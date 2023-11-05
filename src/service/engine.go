package service

import (
	gossr "github.com/natewong1313/go-react-ssr"
)

var APP_ENV string

const title = "Enzo Acad"

func StartEngine() *gossr.Engine {
	engine, err := gossr.New(gossr.Config{
		AppEnv:              "production",
		AssetRoute:          "/assets",
		FrontendDir:         "./frontend/src",
		GeneratedTypesPath:  "./frontend/src/generated.d.ts",
		TailwindConfigPath:  "./frontend/tailwind.config.js",
		LayoutCSSFilePath:   "Main.css",
		PropsStructsPath:    "src/props",
		HotReloadServerPort: 3001,
	})

	if err != nil {
		panic(err)
	}
	return engine
}

var Engine = StartEngine()

func Render(file string, engine *gossr.Engine, prop interface{}) string {
	response := engine.RenderRoute(gossr.RenderConfig{
		File:     file,
		Title:    "Enzo Acad",
		MetaTags: nil,
		Props:    prop,
	})

	return string(response)
}

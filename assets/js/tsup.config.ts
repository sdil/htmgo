import { defineConfig } from "tsup";

export default defineConfig({
  format: ["esm"],
  entry: ["mhtml.ts", "./scripts/*.ts"],
  outDir: "./../dist",
  dts: false,
  shims: true,
  skipNodeModulesBundle: true,
  clean: false,
  target: "esnext",
  treeshake: false,
  platform: "browser",
  outExtension: () => {
    return {
      js: ".js",
    };
  },
  minify: false,
  bundle: true,
  // https://github.com/egoist/tsup/issues/619
  noExternal: [/(.*)/],
});

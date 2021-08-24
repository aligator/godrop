const path = require("path");

module.exports = {
    chainWebpack: config => {
        config
            .entry("app")
            .clear()
            .add("./ui/src/main.ts")
            .end();
        config.resolve.alias
            .set("@", path.join(__dirname, "./ui/src"))
    }
};

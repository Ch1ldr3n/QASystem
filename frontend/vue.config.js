module.exports = {
	devServer: {
		proxy: {
			'/v1': {
				target: "https://qanda-bauhinia.app.secoder.net/",
				changOrigin: true,
				ws: true,
				pathRewrite: {
					'^/v1': ''
				}
			}
		}
	}
  }
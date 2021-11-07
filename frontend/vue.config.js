module.exports = {
  devServer: {
    proxy: {
      '/v1': {
        target: 'http://127.0.0.1:8081',
        changOrigin: true,
        ws: true,
      },
    },
  },
};

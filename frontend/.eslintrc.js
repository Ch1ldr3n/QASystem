module.exports = {
  env: {
    browser: true,
    es2021: false,
  },
  extends: [
    'plugin:vue/essential',
    'airbnb-base',
  ],
  parserOptions: {
    ecmaVersion: 8,
    sourceType: 'module',
  },
  plugins: [
    'vue',
  ],
  rules: {
    "max-len": "off",
    "vue/multi-word-component-names": "off",
  },
};

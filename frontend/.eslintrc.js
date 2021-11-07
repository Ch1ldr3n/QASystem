module.exports = {
  env: {
    browser: true,
    es2021: false,
  },
  extends: [
    'plugin:vue/vue3-recommended',
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
    'max-len': 'off',
    'vue/multi-word-component-names': 'off',
    'linebreak-style': [
      "error",
      "windows"
    ],
  },
  },
};

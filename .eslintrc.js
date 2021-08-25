/* eslint-disable no-undef */
module.exports = {
  parser: '@typescript-eslint/parser', // Specifies the ESLint parser
  plugins: [
    '@typescript-eslint' // add the TypeScript plugin
  ],
  extends: [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    'plugin:@typescript-eslint/recommended-requiring-type-checking'

  ],
  overrides: [
    {
      files: ['*.ts', '*.tsx'], // Your TypeScript files extension
      parserOptions: {
        tsconfigRootDir: __dirname,
        project: ['./tsconfig.json'],
      },
    }
  ],
  rules: {
    "@typescript-eslint/semi": "off",
    "@typescript-eslint/no-unused-vars": [
      "error",
      {
        "varsIgnorePattern": "PlainJSX"
      }
    ]
  },
  ignorePatterns: [
    "types.ts",
  ],
  env: {
    browser: true,
  },
};
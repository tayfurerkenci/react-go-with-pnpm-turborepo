const { configureStore } = require('@reduxjs/toolkit');

module.exports = {
  schemaFile: '../oas/openapi.yaml',
  apiFile: './src/store/baseApi.ts',
  apiImport: 'baseApi',
  outputFile: './src/generated/api.ts',
  exportName: 'moviesApi',
  hooks: true,
  tag: true,
};

if (Test-Path node_modules){(Remove-Item -Recurse -Force node_modules)}
yarn install
npm run dev
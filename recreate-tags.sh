#!/bin/bash

TAG="v0.3.0"

git add .
git commit -m "📦 updates modules for ${TAG}"

git tag ${TAG}
git tag commons/${TAG}
git tag capsulemodule/${TAG}

git push origin main ${TAG} commons/${TAG} capsulemodule/${TAG}


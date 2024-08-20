image_tag=$1
if [ "$image_tag" = "" ]; then
  echo "Usage: $0 <image_tag>"
  exit 0
fi

# build server (golang)
server_image="registry.cn-hangzhou.aliyuncs.com/gx-logistics/gx-customs-bridge:${image_tag}"
docker build -t "${server_image}" -f Dockerfile.server .
# docker buildx build --platform linux/arm64 -t "${server_image}" -f Dockerfile-go -o type=docker .

docker image prune -f
echo "docker push $server_image"
docker push "${server_image}"

echo $TRAVIS_OS_NAME
travis_retry sudo add-apt-repository -y ppa:graphics-drivers/ppa
travis_retry sudo apt-get update -qq
travis_retry sudo apt-get install -f -y nvidia-378 nvidia-378-dev libcuda1-378 nvidia-settings nvidia-opencl-icd-378
travis_retry sudo apt-get install -f

travis_retry wget http://developer.download.nvidia.com/compute/cuda/repos/ubuntu1404/x86_64/cuda-repo-ubuntu1404_${CUDA}_amd64.deb
travis_retry sudo dpkg -i cuda-repo-ubuntu1404_${CUDA}_amd64.deb
travis_retry sudo apt-get update -qq

export CUDA_APT=${CUDA:0:3}
export CUDA_APT=${CUDA_APT/./-}
travis_retry sudo apt-get install -f -y cuda-drivers cuda-core-${CUDA_APT} cuda-cudart-dev-${CUDA_APT} cuda-cufft-dev-${CUDA_APT}
travis_retry sudo apt-get install cuda
travis_retry sudo apt-get clean

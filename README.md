## How to run

```bash
./netprobe -help
Usage of CLI tool:
  -stats: Show network statistics
  -help: Show help information
```

## After Code Change

generate a new tar.gz file with your current code changes

```bash
git add .
git commit -m "Your commit message"
git push origin main
git tag v1.0.2
git push origin v1.0.2
```

GitHub automatically generates tar.gz files for each release of your project. You can find the new file in the "Releases" section of your GitHub repository. The URL of the file will be:

https://github.com/LordMoMA/NetProbe/archive/refs/tags/v1.0.2.tar.gz

## Generate sha256 hash using the above downloaded tarball

Go to the folder where you downloaded the tarball and run the following command:

```bash
shasum -a 256 NetProbe-1.0.2.tar.gz
```

## Update Homebrew package

```bash
brew update
brew upgrade netprobe
```

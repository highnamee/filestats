#!/usr/bin/env bash
set -euo pipefail

VERSION="${1:?usage: scripts/release.sh <version>  e.g. 0.2.0}"
VERSION="${VERSION#v}" # normalise: strip leading 'v' if present

BIN="filestats"
REPO="highnamee/filestats"
LDFLAGS="-s -w -X main.version=${VERSION}"
DIST="dist"

# ── guards ────────────────────────────────────────────────────────────────────
if [[ -n "$(git status --porcelain)" ]]; then
  echo "error: working tree is dirty — commit or stash changes first" >&2
  exit 1
fi

if git rev-parse "v${VERSION}" &>/dev/null; then
  echo "error: tag v${VERSION} already exists" >&2
  exit 1
fi

# ── build ─────────────────────────────────────────────────────────────────────
echo "Building ${BIN} v${VERSION} …"
mkdir -p "${DIST}"

build() {
  local os="$1" arch="$2"
  local out="${DIST}/${BIN}-${os}-${arch}"
  printf "  %-20s → %s\n" "${os}/${arch}" "${out}" >&2
  GOOS="${os}" GOARCH="${arch}" go build -ldflags "${LDFLAGS}" -o "${out}" .
  shasum -a 256 "${out}" | awk '{print $1}'
}

SHA_DARWIN_ARM64="$(build darwin arm64)"
SHA_DARWIN_AMD64="$(build darwin amd64)"
SHA_LINUX_ARM64="$(build linux arm64)"
SHA_LINUX_AMD64="$(build linux amd64)"

# ── update formula ────────────────────────────────────────────────────────────
echo "Updating Formula/filestats.rb …"

cat > Formula/filestats.rb << FORMULA
class Filestats < Formula
  desc "Count file statistics by extension, similar to GitHub's language breakdown"
  homepage "https://github.com/${REPO}"
  version "${VERSION}"

  on_macos do
    on_arm do
      url "https://github.com/${REPO}/releases/download/v#{version}/filestats-darwin-arm64"
      sha256 "${SHA_DARWIN_ARM64}"
    end
    on_intel do
      url "https://github.com/${REPO}/releases/download/v#{version}/filestats-darwin-amd64"
      sha256 "${SHA_DARWIN_AMD64}"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/${REPO}/releases/download/v#{version}/filestats-linux-arm64"
      sha256 "${SHA_LINUX_ARM64}"
    end
    on_intel do
      url "https://github.com/${REPO}/releases/download/v#{version}/filestats-linux-amd64"
      sha256 "${SHA_LINUX_AMD64}"
    end
  end

  def install
    bin.install Dir["filestats-*"].first => "filestats"
  end

  test do
    system "#{bin}/filestats", "-version"
  end
end
FORMULA

# ── tag ───────────────────────────────────────────────────────────────────────
git add Formula/filestats.rb
git commit -m "release: v${VERSION}"
git tag "v${VERSION}"

# ── next steps ────────────────────────────────────────────────────────────────
echo ""
echo "Done. Push and publish:"
echo ""
echo "  git push origin main v${VERSION}"
echo "  gh release create v${VERSION} ${DIST}/${BIN}-* --title \"v${VERSION}\""
echo ""
echo "Then copy Formula/filestats.rb to highnamee/homebrew-filestats and push."

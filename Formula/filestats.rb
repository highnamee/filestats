class Filestats < Formula
  desc "Count file statistics by extension, similar to GitHub's language breakdown"
  homepage "https://github.com/highnamee/filestats"
  version "1.0.0"

  on_macos do
    on_arm do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-darwin-arm64"
      sha256 "76ed3794dff89e4bb58423e649ff6213d99e25dcb0d93117d5585f6284f49245"
    end
    on_intel do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-darwin-amd64"
      sha256 "82f48cea4644c443e6780967366ee61091ba41bb8810d5a5caf0e57fb4394f8b"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-linux-arm64"
      sha256 "eedd61b8d8f10d499bd242d5a7338bbb2620af771f9b3ff26d55b4488670aae6"
    end
    on_intel do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-linux-amd64"
      sha256 "e8f7f76fb6c7c67f8713883cbc0144f21aa29eedc3e20ca5938f4d44afbecfed"
    end
  end

  def install
    bin.install Dir["filestats-*"].first => "filestats"
  end

  test do
    system "#{bin}/filestats", "-version"
  end
end

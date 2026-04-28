class Filestats < Formula
  desc "Count file statistics by extension, similar to GitHub's language breakdown"
  homepage "https://github.com/highnamee/filestats"
  version "1.3.0"

  on_macos do
    on_arm do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-darwin-arm64"
      sha256 "ad89705529c469a1dfbd1be5fefc1ffb6feb355b1d2d84f281653234871912fd"
    end
    on_intel do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-darwin-amd64"
      sha256 "0154c10bc6fe94b3307651b424b94445d3b9cf2543cef811a14d37924d60e09c"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-linux-arm64"
      sha256 "fd8dbde33dba2eee4497db0012efe3c3c014b1c5f871d5038b1d30ffaf190b6c"
    end
    on_intel do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-linux-amd64"
      sha256 "5b258e33b775cf2f823951b6440ada4a0bb8308bc6afff43ed39f86dde64356b"
    end
  end

  def install
    bin.install Dir["filestats-*"].first => "filestats"
  end

  test do
    system "#{bin}/filestats", "-version"
  end
end

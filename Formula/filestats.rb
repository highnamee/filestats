class Filestats < Formula
  desc "Count file statistics by extension, similar to GitHub's language breakdown"
  homepage "https://github.com/highnamee/filestats"
  version "1.1.0"

  on_macos do
    on_arm do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-darwin-arm64"
      sha256 "d8b5b157c45c3c52f761b5f3f5ef25fca8783c6ffd542ca0801bcd002a219adc"
    end
    on_intel do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-darwin-amd64"
      sha256 "7a6071b12e1846ac38b2eee1717ac403a51e12f408b68b07d7b138dbebfc1a2b"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-linux-arm64"
      sha256 "b4f732576f252527929e45813b82c6c8a3c803489738349ca856b919209de6a0"
    end
    on_intel do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-linux-amd64"
      sha256 "02c64bf8dd6723e8c319119aa47a09e7d494103fa07a94ff73f5d7d5921d5b58"
    end
  end

  def install
    bin.install Dir["filestats-*"].first => "filestats"
  end

  test do
    system "#{bin}/filestats", "-version"
  end
end

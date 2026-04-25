class Filestats < Formula
  desc "Count file statistics by extension, similar to GitHub's language breakdown"
  homepage "https://github.com/highnamee/filestats"
  version "1.1.1"

  on_macos do
    on_arm do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-darwin-arm64"
      sha256 "5611e9254ce1acfc7987c0fef05413dc32f63b07ed3aca00e2e6ce2e27d33963"
    end
    on_intel do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-darwin-amd64"
      sha256 "2ff11dc219ef04b54ef34e6b8a3b93b589e8101f5141fe88a758850bc4735047"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-linux-arm64"
      sha256 "f4a19ad5ab99308c2544e87837c52e8bc1e885c868865821cc033f3b6a9859e0"
    end
    on_intel do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-linux-amd64"
      sha256 "17bd2b240b854109125fa311b2430fff3694deac593cd036adf102b6158b9448"
    end
  end

  def install
    bin.install Dir["filestats-*"].first => "filestats"
  end

  test do
    system "#{bin}/filestats", "-version"
  end
end

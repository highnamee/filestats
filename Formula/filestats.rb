class Filestats < Formula
  desc "Count file statistics by extension, similar to GitHub's language breakdown"
  homepage "https://github.com/highnamee/filestats"
  version "1.2.0"

  on_macos do
    on_arm do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-darwin-arm64"
      sha256 "b50a8e16c2dc527ea306ecc9525d45daab4c12293f5b04576d3f3745c1953580"
    end
    on_intel do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-darwin-amd64"
      sha256 "79ba3e2dd2013a061fd02ada417de7f3a8e96b1d87ea90d4b47829d1a592436f"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-linux-arm64"
      sha256 "dc2b5db8f8de05f87c412090339b0d38ca1d627659d06bfdd75ccc4dbc0f392a"
    end
    on_intel do
      url "https://github.com/highnamee/filestats/releases/download/v#{version}/filestats-linux-amd64"
      sha256 "1fb3762aa31469a999e5af0aafd17ac8be5eb69392a67ec9f83c4ffb7b08579e"
    end
  end

  def install
    bin.install Dir["filestats-*"].first => "filestats"
  end

  test do
    system "#{bin}/filestats", "-version"
  end
end

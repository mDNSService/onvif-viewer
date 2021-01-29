class onvif-viewer < Formula
  desc "Onvif Viewer"
  homepage "https://github.com/mDNSService/onvif-viewer"
  url "https://github.com/mDNSService/onvif-viewer.git",
      tag:      "v0.0.5",
      revision: "bebda34d7f36148bc410706063fb9cec6d9ff1df"
  license "MIT"

  depends_on "go" => :build

  def install
    (etc/"onvif-viewer").mkpath
    system "go", "build", "-mod=vendor", "-ldflags",
             "-s -w -X main.version=#{version} -X main.commit=#{stable.specs[:revision]} -X main.builtBy=homebrew",
             *std_go_args
    etc.install "onvif-viewer.yaml" => "onvif-viewer/onvif-viewer.yaml"
  end

  plist_options manual: "onvif-viewer -c #{HOMEBREW_PREFIX}/etc/onvif-viewer/onvif-viewer.yaml"

  def plist
    <<~EOS
      <?xml version="1.0" encoding="UTF-8"?>
      <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
      <plist version="1.0">
        <dict>
          <key>Label</key>
          <string>#{plist_name}</string>
          <key>KeepAlive</key>
          <true/>
          <key>ProgramArguments</key>
          <array>
            <string>#{opt_bin}/onvif-viewer</string>
            <string>-c</string>
            <string>#{etc}/onvif-viewer/onvif-viewer.yaml</string>
          </array>
          <key>StandardErrorPath</key>
          <string>#{var}/log/onvif-viewer.log</string>
          <key>StandardOutPath</key>
          <string>#{var}/log/onvif-viewer.log</string>
        </dict>
      </plist>
    EOS
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/onvif-viewer -v 2>&1")
    assert_match "config created", shell_output("#{bin}/onvif-viewer init --config=onvif-viewer.yml 2>&1")
    assert_predicate testpath/"onvif-viewer.yml", :exist?
  end
end

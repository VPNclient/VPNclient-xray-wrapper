# 🚀 VPN Client XRAY Wrapper ('vpnclient_xray_wrapper')

A high-performance VPN client wrapper for Xray-core, designed to provide secure and private internet access. This project implements VPN functionality for Android, iOS, macOS, Windows and Linux platforms.

## 🌟 Features
- 🔒 Secure VPN tunneling using Xray-core  
- 🔄 Support for multiple protocols (VMess, VLESS, Trojan, etc.)  
- 📱 Cross-platform implementation (Android/iOS)  
- ⚡ Optimized for performance and battery efficiency  
- 🛠️ Configurable through simple JSON configuration  

## 📋 Requirements
- 🤖 Android 6.0+, 🍎 iOS 12.0+  
- ✅ Appropriate VPN permissions  

## 📥 Installation

### 🤖 Android
1. 📂 Clone this repository  
2. 🚀 Open the project in Android Studio  
3. 🔨 Build and run the application  

```bash
git clone https://github.com/VPNclient/VPNclient-xray-wrapper.git
cd VPNclient-xray-wrapper/android
./gradlew assembleDebug
```

### 🍎 iOS
1. 📂 Clone this repository  
2. 🚀 Open the project in Xcode  
3. 🔨 Build and run the application  

```bash
git clone https://github.com/VPNclient-repo/VPNclient-xray-wrapper.git
open VPNclient-xray-wrapper/ios/VPNClient.xcodeproj
```

## ⚙️ Configuration
Create a `config.json` file with your Xray-core configuration:

```json
{
  "inbounds": [...],
  "outbounds": [...],
  "routing": {...}
}
```

## 🚀 Usage
1. 📁 Place your Xray configuration in the app's documents directory  
2. ▶️ Launch the application  
3. 📑 Select your configuration file  
4. 🔗 Tap "Connect" to establish the VPN connection  


## 🤝 Contributing
We welcome contributions! Please fork the repository and submit pull requests.

## 📜 License

This project is licensed under the **VPNclient Extended GNU General Public License v3 (GPL v3)**. See [LICENSE.md](LICENSE.md) for details.

⚠️ **Note:** By using this software, you agree to comply with additional conditions outlined in the [VPNсlient Extended GNU General Public License v3 (GPL v3)](LICENSE.md)




## 💬 Support
For issues or questions, please open an issue on our GitHub repository.

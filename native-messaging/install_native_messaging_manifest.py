import json
import os.path
import platform

NAME = "native_messaging_example_app"

REG_ENTRY = {
    "name": NAME,
    "description": "Example host for native messaging",
    "path": os.path.join(os.path.dirname(os.path.abspath(__file__)), "dist", "app_bin"),
    "type": "stdio",
    "allowed_extensions": ["native_messaging_example@tfinlay.io"]
}


if __name__ == "__main__":
    if platform.system() == "Windows":
        import winreg

        config_path = os.path.join(os.path.dirname(os.path.abspath(__file__)), f"dist\\{NAME}.json")
        with open(config_path, 'w') as f:
            json.dump(REG_ENTRY, f, indent=4)

        KEY_PATH = f"Software\\Mozilla\\NativeMessagingHosts\\{NAME}"
        key = winreg.CreateKey(winreg.HKEY_CURRENT_USER, KEY_PATH)
        winreg.SetValueEx(key, '', 0, winreg.REG_SZ, config_path)
    elif platform.system() == "Darwin":
        with open(f'~/Library/Application Support/Mozilla/NativeMessagingHosts/{NAME}.json"', 'w') as f:
            json.dump(REG_ENTRY, f, indent=4)
    else:
        raise NotImplementedError(f"System {platform.system()} is not suported")

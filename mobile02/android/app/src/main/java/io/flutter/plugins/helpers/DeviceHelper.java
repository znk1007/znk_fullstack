package io.flutter.plugins.helpers;

import java.util.List;

import android.content.Context;

import android.content.ContentResolver;
import androidx.annotation.NonNull;
import io.flutter.plugin.common.BinaryMessenger;
import io.flutter.plugin.common.MethodChannel;
import io.flutter.plugin.common.PluginRegistry.Registrar;
import io.flutter.embedding.engine.FlutterEngine;
import io.flutter.embedding.engine.plugins.FlutterPlugin;
import io.flutter.plugin.common.MethodCall;
import io.flutter.plugin.common.EventChannel;

import android.annotation.SuppressLint;
import android.os.Build;
import android.content.ContentResolver;
import android.provider.Settings;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

@SuppressWarnings("unchecked")
public class DeviceHelper implements FlutterPlugin {

    MethodChannel methodChannel;

    private Context context;


    public static void registerWith(Registrar registrar) {
        // helper
        DeviceHelper helper = new DeviceHelper();
        helper.context = registrar.context();
        helper.setupMethodChannel(
            registrar.messenger(), 
            registrar.context().getContentResolver()
        );
    }

    @Override
    public void onAttachedToEngine(FlutterPlugin.FlutterPluginBinding binding) {
        context = binding.getApplicationContext();
        setupMethodChannel(
            binding.getFlutterEngine().getDartExecutor(),
            binding.getApplicationContext().getContentResolver()
        );
    }

    @Override
    public void onDetachedFromEngine(FlutterPlugin.FlutterPluginBinding binding) {
        teardownMethodChannel();
    }


    private void setupMethodChannel(BinaryMessenger messenger, ContentResolver contentResolver) {
        methodChannel = new MethodChannel(messenger, "device_helper_channel");
        final DeviceHelperImpl handler = new DeviceHelperImpl(contentResolver);
        methodChannel.setMethodCallHandler(handler);
    }

    private void teardownMethodChannel() {
        methodChannel.setMethodCallHandler(null);
        methodChannel = null;
    }
    
}

class DeviceHelperImpl implements MethodChannel.MethodCallHandler {

    private ContentResolver contentResolver;

    private static final String[] EMPTY_STRING_LIST = new String[] {};
    
    /** Constructs MethodCallHandlerImpl. The {@code contentResolver} must not be null. */
    DeviceHelperImpl(ContentResolver contentResolver) {
        this.contentResolver = contentResolver;
    }

    @Override
    public void onMethodCall(MethodCall call, MethodChannel.Result result) {
        switch (call.method) {
            case "getAndroidDeviceInfo":
            {
                getDeviceInfoWithResult(result);
            }
                break;
            default:
            {
                result.notImplemented();
            }
                break;
        }
    }

    private void getDeviceInfoWithResult(MethodChannel.Result result) {
        Map<String, Object> build = new HashMap<>();
        build.put("board", Build.BOARD);
        build.put("bootloader", Build.BOOTLOADER);
        build.put("brand", Build.BRAND);
        build.put("device", Build.DEVICE);
        build.put("display", Build.DISPLAY);
        build.put("fingerprint", Build.FINGERPRINT);
        build.put("hardware", Build.HARDWARE);
        build.put("host", Build.HOST);
        build.put("id", Build.ID);
        build.put("manufacturer", Build.MANUFACTURER);
        build.put("model", Build.MODEL);
        build.put("product", Build.PRODUCT);
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.LOLLIPOP) {
            build.put("supported32BitAbis", Arrays.asList(Build.SUPPORTED_32_BIT_ABIS));
            build.put("supported64BitAbis", Arrays.asList(Build.SUPPORTED_64_BIT_ABIS));
            build.put("supportedAbis", Arrays.asList(Build.SUPPORTED_ABIS));
        } else {
            build.put("supported32BitAbis", Arrays.asList(EMPTY_STRING_LIST));
            build.put("supported64BitAbis", Arrays.asList(EMPTY_STRING_LIST));
            build.put("supportedAbis", Arrays.asList(EMPTY_STRING_LIST));
        }
        build.put("tags", Build.TAGS);
        build.put("type", Build.TYPE);
        build.put("isPhysicalDevice", !isEmulator());
        build.put("androidId", getAndroidId());

        Map<String, Object> version = new HashMap<>();
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) {
            version.put("baseOS", Build.VERSION.BASE_OS);
            version.put("previewSdkInt", Build.VERSION.PREVIEW_SDK_INT);
            version.put("securityPatch", Build.VERSION.SECURITY_PATCH);
        }
        version.put("codename", Build.VERSION.CODENAME);
        version.put("incremental", Build.VERSION.INCREMENTAL);
        version.put("release", Build.VERSION.RELEASE);
        version.put("sdkInt", Build.VERSION.SDK_INT);
        build.put("version", version);
        result.success(build);
    }

    @SuppressLint("HardwareIds")
    private String getAndroidId() {
        return Settings.Secure.getString(contentResolver, Settings.Secure.ANDROID_ID);
    }

    
    private boolean isEmulator() {
        return (Build.BRAND.startsWith("generic") && Build.DEVICE.startsWith("generic"))
            || Build.FINGERPRINT.startsWith("generic")
            || Build.FINGERPRINT.startsWith("unknown")
            || Build.HARDWARE.contains("goldfish")
            || Build.HARDWARE.contains("ranchu")
            || Build.MODEL.contains("google_sdk")
            || Build.MODEL.contains("Emulator")
            || Build.MODEL.contains("Android SDK built for x86")
            || Build.MANUFACTURER.contains("Genymotion")
            || Build.PRODUCT.contains("sdk_google")
            || Build.PRODUCT.contains("google_sdk")
            || Build.PRODUCT.contains("sdk")
            || Build.PRODUCT.contains("sdk_x86")
            || Build.PRODUCT.contains("vbox86p")
            || Build.PRODUCT.contains("emulator")
            || Build.PRODUCT.contains("simulator");
    }
}
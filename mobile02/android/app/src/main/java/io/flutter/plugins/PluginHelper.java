package io.flutter.plugins;

import java.util.List;

import android.content.ContentResolver;
import androidx.annotation.NonNull;
import io.flutter.plugin.common.BinaryMessenger;
import io.flutter.plugin.common.MethodChannel;
import io.flutter.plugin.common.PluginRegistry.Registrar;
import io.flutter.embedding.engine.FlutterEngine;
import io.flutter.embedding.engine.plugins.FlutterPlugin;
import io.flutter.plugin.common.MethodCall;
import io.flutter.plugin.common.EventChannel;
import io.flutter.plugins.helpers.DeviceHelper;

@SuppressWarnings("unchecked")
public class PluginHelper implements FlutterPlugin {

    MethodChannel methodChannel;
    EventChannel eventChannel;


    // public static void registerWith(Registrar registrar) {
    //     // helper
    //     PluginHelper helper = new PluginHelper();
    //     helper.setupMethodChannel(
    //         registrar.messenger(), 
    //         registrar.context().getContentResolver()
    //     );
    // }

    @Override
    public void onAttachedToEngine(FlutterPlugin.FlutterPluginBinding binding) {
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
        methodChannel = new MethodChannel(messenger, "method_channel_helper");
        final MethodCallHandlerImpl handler = new MethodCallHandlerImpl(contentResolver);
        methodChannel.setMethodCallHandler(handler);
    }

    private void teardownMethodChannel() {
        methodChannel.setMethodCallHandler(null);
        methodChannel = null;
    }
    
}

class MethodCallHandlerImpl implements MethodChannel.MethodCallHandler {

    private ContentResolver contentResolver;
    
    /** Constructs MethodCallHandlerImpl. The {@code contentResolver} must not be null. */
    MethodCallHandlerImpl(ContentResolver contentResolver) {
        this.contentResolver = contentResolver;
    }

    @Override
    public void onMethodCall(MethodCall call, MethodChannel.Result result) {
        if (call.method.equals("getAndroidDeviceInfo")) {
            DeviceHelper device = new DeviceHelper(contentResolver);
            device.addDeviceInfo();
            result.success(device.build);
        } else {
            result.notImplemented();
        }
    }
}
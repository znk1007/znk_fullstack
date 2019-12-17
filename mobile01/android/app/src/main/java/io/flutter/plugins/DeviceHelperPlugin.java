package io.flutter.plugins;

import android.content.ContentResolver;
import io.flutter.plugin.embedding.engine.plugins.FlutterPlugin;
import io.flutter.plugin.common.BinaryMessenger;
import io.flutter.plugin.common.MethodChannel;
import io.flutter.plugin.common.PluginRegistry.Registrar;

public class DeviceHelperPlugin implements FlutterPlugin {

    MethodChannel channel;

    public static void registerWith(Registrar registrar) {
        DeviceHelperPlugin plugin = new DeviceHelperPlugin();
        plugin.setupMethodChannel(registrar.messenger(), registrar.context().getContentResolver());
    }

    @Override
    public void onAttackedToEngine(FlutterPlugin.FlutterPluginBinding binding) {
        setupMethodChannel(
            binding.getFlutterEngine().getDartExecutor(), 
            binding.getApplicationContext().getContentResolver()
        );
    }

    @Override
    public void onDetachedFromEngine(FlutterPlugin.FlutterPluginBinding binding) {
        tearDownChannel();
    }
    

    private void setupMethodChannel(BinaryMessenger messenger, ContentResolver contentResolver) {
        channel = new MethodChannel(messenger, "device_helper");
        final MethodCallHandlerImpl handler = new MethodCallHandlerImpl(contentResolver);
        channel.setMethodCallHandler(handler);
    }

    private void tearDownChannel() {
        channel.setMethodCallHandler(null);
        channel = null;
    }
}
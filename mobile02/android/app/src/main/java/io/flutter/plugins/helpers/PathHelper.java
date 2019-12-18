package io.flutter.plugins.helpers;
import java.util.List;
import java.util.logging.Logger;

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
import android.os.Build.VERSION;
import android.os.Build.VERSION_CODES;
import androidx.annotation.NonNull;
import android.os.Environment;
import io.flutter.util.PathUtils;
import java.io.File;
import java.util.ArrayList;
import java.util.List;

public class PathHelper implements FlutterPlugin, MethodChannel.MethodCallHandler {

    private Context context;
    private MethodChannel channel;
  
    public PathHelper() {}
  
    public static void registerWith(Registrar registrar) {
        PathHelper instance = new PathHelper();
        instance.channel = new MethodChannel(registrar.messenger(), "path_helper_channel");
        instance.context = registrar.context();
        instance.channel.setMethodCallHandler(instance);
    }
  
    @Override
    public void onAttachedToEngine(@NonNull FlutterPluginBinding binding) {
        channel =
            new MethodChannel(
                binding.getFlutterEngine().getDartExecutor(), "path_helper_channel");
        context = binding.getApplicationContext();
        channel.setMethodCallHandler(this);
    }
  
    @Override
    public void onDetachedFromEngine(@NonNull FlutterPluginBinding binding) {
      channel.setMethodCallHandler(null);
      channel = null;
    }
  
    @Override
    public void onMethodCall(MethodCall call, @NonNull MethodChannel.Result result) {
      switch (call.method) {
        case "getTemporaryDirectory":
          result.success(getPathProviderTemporaryDirectory());
          break;
        case "getApplicationDocumentsDirectory":
          result.success(getPathProviderApplicationDocumentsDirectory());
          break;
        case "getStorageDirectory":
          result.success(getPathProviderStorageDirectory());
          break;
        case "getExternalCacheDirectories":
          result.success(getPathProviderExternalCacheDirectories());
          break;
        case "getExternalStorageDirectories":
          final Integer type = call.argument("type");
          final String directoryName = androidType(type);
          result.success(getPathProviderExternalStorageDirectories(directoryName));
          break;
        case "getApplicationSupportDirectory":
          result.success(getApplicationSupportDirectory());
          break;
        default:
          result.notImplemented();
      }
    }
  
    private String getPathProviderTemporaryDirectory() {
      return context.getCacheDir().getPath();
    }
  
    private String getApplicationSupportDirectory() {
      return PathUtils.getFilesDir(context);
    }
  
    private String getPathProviderApplicationDocumentsDirectory() {
      return PathUtils.getDataDirectory(context);
    }
  
    private String getPathProviderStorageDirectory() {
      final File dir = context.getExternalFilesDir(null);
      if (dir == null) {
        return null;
      }
      return dir.getAbsolutePath();
    }
  
    private List<String> getPathProviderExternalCacheDirectories() {
      final List<String> paths = new ArrayList<>();
  
      if (VERSION.SDK_INT >= VERSION_CODES.KITKAT) {
        for (File dir : context.getExternalCacheDirs()) {
          if (dir != null) {
            paths.add(dir.getAbsolutePath());
          }
        }
      } else {
        File dir = context.getExternalCacheDir();
        if (dir != null) {
          paths.add(dir.getAbsolutePath());
        }
      }
  
      return paths;
    }
  
    private List<String> getPathProviderExternalStorageDirectories(String type) {
      final List<String> paths = new ArrayList<>();
  
      if (VERSION.SDK_INT >= VERSION_CODES.KITKAT) {
        for (File dir : context.getExternalFilesDirs(type)) {
          if (dir != null) {
            paths.add(dir.getAbsolutePath());
          }
        }
      } else {
        File dir = context.getExternalFilesDir(type);
        if (dir != null) {
          paths.add(dir.getAbsolutePath());
        }
      }
  
      return paths;
    }

    /**
     * Return a Android Environment constant for a Dart Index.
     *
     * @return The correct Android Environment constant or null, if the index is null.
     * @throws IllegalArgumentException If `dartIndex` is not null but also not matches any known
     *     index.
     */
    String androidType(Integer dartIndex) throws IllegalArgumentException {
        if (dartIndex == null) {
            return null;
        }

        switch (dartIndex) {
        case 0:
            return Environment.DIRECTORY_MUSIC;
        case 1:
            return Environment.DIRECTORY_PODCASTS;
        case 2:
            return Environment.DIRECTORY_RINGTONES;
        case 3:
            return Environment.DIRECTORY_ALARMS;
        case 4:
            return Environment.DIRECTORY_NOTIFICATIONS;
        case 5:
            return Environment.DIRECTORY_PICTURES;
        case 6:
            return Environment.DIRECTORY_MOVIES;
        case 7:
            return Environment.DIRECTORY_DOWNLOADS;
        case 8:
            return Environment.DIRECTORY_DCIM;
        case 9:
            if (VERSION.SDK_INT >= VERSION_CODES.KITKAT) {
            return Environment.DIRECTORY_DOCUMENTS;
            } else {
            throw new IllegalArgumentException("Documents directory is unsupported.");
            }
        default:
            throw new IllegalArgumentException("Unknown index: " + dartIndex);
        }
    }

  }


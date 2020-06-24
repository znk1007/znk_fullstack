//
//  NSObject+Ext.swift
//  FLBMall
//
//  Created by Sam Huang on 2020/6/24.
//  Copyright © 2020 Sam Huang. All rights reserved.
//

import UIKit

extension UIColor {
    
    /// 随机色
    static var randomColor: UIColor {
        get {
            let red = CGFloat(arc4random()%256)/255.0
            let green = CGFloat(arc4random()%256)/255.0
            let blue = CGFloat(arc4random()%256)/255.0
            return UIColor(red: red, green: green, blue: blue, alpha: 1.0)
        }
    }
}

//
//  ViewController.swift
//  FLBMall
//
//  Created by Sam Huang on 2020/6/24.
//  Copyright © 2020 Sam Huang. All rights reserved.
//

import UIKit

class ViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
//        self.view.backgroundColor = UIColor.blue
        
        
    
    }

    override func viewWillAppear(_ animated: Bool) {
        super.viewWillAppear(animated)
        let gradientLayer = CAGradientLayer.init()
        gradientLayer.frame = self.view.bounds
        gradientLayer.colors = [UIColor.red.cgColor, UIColor.green.cgColor, UIColor.yellow.cgColor]
        gradientLayer.locations = [0.0, 0.05, 0.2]
        
        gradientLayer.startPoint = CGPoint.init(x: 0, y: 0)
        gradientLayer.endPoint = CGPoint.init(x: 0, y: 1)
        self.view.layer.addSublayer(gradientLayer)
    }
}


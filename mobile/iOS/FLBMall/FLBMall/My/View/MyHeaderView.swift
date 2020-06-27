//
//  MyHeaderView.swift
//  FLBMall
//
//  Created by Sam Huang on 2020/6/24.
//  Copyright © 2020 Sam Huang. All rights reserved.
//

import UIKit

enum HeaderExtraType {
    case rate//积分
    case packet//红包
}

fileprivate class HeaderExtraView: UIView {
    
    
    
    lazy var titleLab: UILabel = {
        $0.text = "66"
        $0.textColor = UIColor.black
        $0.font = UIFont.systemFont(ofSize: 12)
        return  $0
    }(UILabel.init())
    
    lazy var subtitleLab: UILabel = {
           $0.text = "我的积分"
           $0.textColor = UIColor.black
           $0.font = UIFont.systemFont(ofSize: 14)
           return  $0
           return  $0
       }(UILabel.init())
    
    lazy var btn: UIButton = {
        $0.addTarget(self, action: #selector(extraItemClick), for: .touchUpInside)
        return $0
    }(UIButton.init(type: .custom))
    
    @objc func extraItemClick() {
        
    }
}

class MyHeaderView: UIView {
    
    fileprivate lazy var extraHeader: HeaderExtraView = {
        
        return $0
    }(HeaderExtraView.init())
    
    var type: HeaderExtraType = .rate {
        didSet {
            switch type {
            case .rate:
                self.extraHeader.titleLab.text = "66"
                self.extraHeader.subtitleLab.text = "我的积分"
            case .packet:
                self.extraHeader.titleLab.text = "88"
                self.extraHeader.subtitleLab.text = "我的红包"
            }
        }
    }

    private lazy var bgView: UIView = {
           $0.backgroundColor = UIColor.randomColor
           return $0
       }(UIView.init())
       
      private lazy var protoBtn: UIButton = {
           $0.layer.masksToBounds = true
           $0.backgroundColor = UIColor.randomColor
           return $0
      }(UIButton.init(type: .custom))
       
      private lazy var nicknameLab: UILabel = {
           $0.text = "昵称"
           $0.textColor = UIColor.white
           $0.backgroundColor = UIColor.randomColor
           return $0
       }(UILabel.init())
       
      private lazy var companyBtn: UIButton = {
           $0 .setTitle("公司信息", for: .normal)
           $0.backgroundColor = UIColor.randomColor
           return $0
       }(UIButton.init(type: .custom))
       
      private lazy var extraInfoView: UIView = {
           $0.layer.masksToBounds = true;
           $0.backgroundColor = UIColor.randomColor
           return $0
       }(UIView.init())


}
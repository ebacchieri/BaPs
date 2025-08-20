![introduce](https://socialify.git.ci/gucooing/BaPs/image?description=1&font=Source+Code+Pro&forks=1&issues=1&language=1&name=1&owner=1&pattern=Plus&pulls=1&stargazers=1&theme=Light)

# BaPs 🎮

#### [Chinese](README.md)

## ⚠️ WARNING: THE ORIGINAL REPOSITORY HAS A REMOTE ACCESS BACKDOOR ⚠️️

The original repository has a developer-only build key **remote shell access backdoor** since [this commit](https://github.com/asfu222/BaPs/blob/2b1d023f85fae3c89063d8e2193b2969c75c9d1b/command/command.go), which was since June 9, 3:51 PM UTC, 2025.

The build key `mx.Key` is set during the private code in package `github.com/ebacchieri/BaPs/protocol/mx` during the build process.

It has had a backdoor for registering bot accounts since [this commit](https://github.com/ebacchieri/BaPs/commit/7e5c80007454bcaeb35d3ad0ead36178f1816bd2#diff-05f9f1d71ee40dc32c7024b65a71f0f355eaede479b769eed8b9e2084ac64f42), and has added an API backdoor since [this commit](https://github.com/ebacchieri/BaPs/blob/388d29902f37bd493a4f9d34208231fbf40b26b1/command/command.go).

See the following screenshots of code snippets as proof for the remote shell access backdoor:

![1](README_media/1.png)

![2](README_media/2.png)

![3](README_media/3.png)

The original developer could've used API posts to `{your_server_ip}/cdq/api/shell` with any arbitrary shell command as query, and that code will run on your machine. The original author called this a "vulnerability" that does not exist in his repository when this fork was presented as an option due to imminent privatization of the repository, and has thrown terms that he does not take any legal responsibility for anything that happens to users of this software. 

That is not legal tenure. It is not even legal at all. In fact, it is **ILLEGAL** in **many if not ALL juristications**, including the **People's Republic of China**.
It went against the following laws in the Chinese criminal code:
- 第二百八十五条 - 中华人民共和国刑法
  > 违反国家规定，侵入国家事务、国防建设、尖端科学技术领域的计算机信息系统的，处三年以下有期徒刑或者拘役。

  > 违反国家规定，侵入前款规定以外的计算机信息系统或者采用其他技术手段，获取该计算机信息系统中存储、处理或者传输的数据，或者对该计算机信息系统实施非法控制，情节严重的，处三年以下有期徒刑或者拘役，并处或者单处罚金；情节特别严重的，处三年以上七年以下有期徒刑，并处罚金。

  > 提供专门用于侵入、非法控制计算机信息系统的程序、工具，或者明知他人实施侵入、非法控制计算机信息系统的违法犯罪行为而为其提供程序、工具，情节严重的，依照前款的规定处罚。{刑法修正案（七）增加第二款、第三款}

- Article 285 — Criminal Law of the People’s Republic of China
  > Whoever, in violation of state regulations, intrudes into computer information systems used for state affairs, national defense construction, or cutting-edge science and technology, shall be sentenced to fixed-term imprisonment of not more than three years or criminal detention.
  > Whoever, in violation of state regulations, intrudes into computer information systems other than those specified above, or uses other technical means to obtain data stored, processed, or transmitted in such systems, or to illegally control such systems, and where the circumstances are serious, shall be sentenced to fixed-term imprisonment of not more than three years, criminal detention, and may also or solely be fined; where the circumstances are especially serious, shall be sentenced to fixed-term imprisonment of not less than three years but not more than seven years, and shall also be fined.
  > Whoever provides programs or tools specifically designed for intruding into or illegally controlling computer information systems, or knowingly provides such programs or tools to others engaged in illegal intrusion or control of computer information systems, where the circumstances are serious, shall be punished in accordance with the preceding paragraph.

- 第二百八十六条 - 中华人民共和国刑法
  > 违反国家规定，对计算机信息系统功能进行删除、修改、增加、干扰，造成计算机信息系统不能正常运行，后果严重的，处五年以下有期徒刑或者拘役；后果特别严重的，处五年以上有期徒刑。

  > 违反国家规定，对计算机信息系统中存储、处理或者传输的数据和应用程序进行删除、修改、增加的操作，后果严重的，依照前款的规定处罚。

  > 故意制作、传播计算机病毒等破坏性程序，影响计算机系统正常运行，后果严重的，依照第一款的规定处罚。

- Article 286 — Criminal Law of the People’s Republic of China
  > Whoever, in violation of state regulations, deletes, modifies, adds to, or interferes with the functions of a computer information system, thereby causing the system to be unable to function normally, and where the consequences are serious, shall be sentenced to fixed-term imprisonment of not more than five years or criminal detention; where the consequences are especially serious, shall be sentenced to fixed-term imprisonment of not less than five years.
  > Whoever, in violation of state regulations, deletes, modifies, or adds data or application programs stored, processed, or transmitted in a computer information system, and where the consequences are serious, shall be punished in accordance with the preceding paragraph.

This is incredibly malicious, and is part of the reason why this fork now exists.
Relevant code snippets that the screenshots were taken from:
1. https://github.com/ebacchieri/BaPs/blob/fd9ce75c83f287022c71e9edb228cae210b7e0b7/command/command.go
2. https://github.com/gucooing/cdq/blob/57ff61f0f476378761ffa70f31d818179ea7a168/api.gin.go
3. https://github.com/gucooing/cdq/blob/57ff61f0f476378761ffa70f31d818179ea7a168/shell.go
#### This fork is a fork of the original project with various patches in the spirit of open-source
#### This fork is made under the platform license grant to all GitHub users by public repositories, relevant legal information [GitHub Terms of Service §D.5](https://docs.github.com/en/site-policy/github-terms/github-terms-of-service#5-license-grant-to-other-users)

#### The original project has major problems with the legality of its "terms of service" or license. It contains clauses that are not enforceable in a court of law, but rather only by blackmail, public shaming, or doxxing. Specifically:
  - > 严禁在中国大陆地区的任何公共或私有平台上传、分享或宣传本项目的源码、编译文件、部署教程、截图、演示视频等相关内容；
    > 若发现违规传播行为，作者有权采取封禁、公开拉黑等措施。

  - This above clause stopping the advertising (and distribution) of the source code is largely unenforceable, and the specified actions likely falls under fair use. As laid out in the platform agreement for public repositories (TOS §D.5), non-exclusive, unrevokable rights to distribute the code using GitHub features are granted to all users. Simply sharing some public GitHub link does not constitute a violation of applicable Fair Use law. Content generated by users, such as tutorials, further fall under the Fair Use category. The author has no legal ground, such as Non-Disclosure Agreements (NDAs), that allows the author to stop original work produced from other users. Thus, such clauses are null and void in a court of law. However, the author does have one part of the clause active, the part that stops the distribution (by uploading) of binaries and source code on Chinese platforms. As such, this fork will not distribute binaries or anything of the kind otherwise due to legality issues. Users are urged to build from source.

- > 禁止开服
严禁使用本项目进行任何形式的私服搭建，包括但不限于公网开服、内网搭建、测试服部署；
  > 无论是否收费、是否开放给他人使用，均视为违反本协议；
  > 一经发现，作者将永久停止支持，并可能公开违规行为及其责任人信息。

- It is unclear the legality of this clause under commercial use. However, this clause is unenforceable and is null and void in a court of law for users that are using it on a personal, noncommercial scale. The clause is only enforceable by blackmail, threatening a maintainance/support stop and public shaming/doxxing. Doxxing is, of course, against the law, both in the juristiction of China and the United States. Applicable violations of Chinese law include 非法获取、使用、传播个人信息 (Illegal use of personal information). No contract is valid legal tenure if its terms do not obey law.

- > 协议更新
  > 本协议可随时更新，使用即视为接受协议的所有内容及后续变更。

- One-sided license changes are not enforceable unless explictly agreed upon. The vague clauses that do not even come in a dedicated license file are not proper notice or legal agreements, and thus are null and void in a court of law.


#### Due to various issues, mostly the original author's issue with people being able to access and run information from the original public project easily. As such, the original author has used various forms of encryption and data gatekeeping to gatekeep builds. (Specifically, there is special data encryption and other processing functions such as mx.DeExcelBytes, which are private and not present in the source code. This was not present at the start of the project)

## Licensing Notice

This repository is a fork of an "All Rights Reserved" project that was hosted publicly on GitHub.

All original code remains under its original restrictions.

However, all modifications and additions made by this fork are licensed under the GNU Affero General Public License v3.0 (AGPL-3.0).

Use of this repository is subject to the terms of the original repository and GitHub Terms of Service §D.5, which allows public forks and modifications within the GitHub platform, only.

### Terms the original project was under
> ⚠️ This project is for learning purposes only, strictly prohibited for commercial use, please delete within 24 hours.

# For learning purposes only, strictly prohibited for commercial use, please delete within 24 hours!!!

> 🌟 Due to its stateless design, it may require slightly higher memory
  
> 📅 Currently supported version: Japan

## 📍Discord

#### Do not contact this discord for support regarding this fork. This fork is unaffliated with the original author, and you shouldn't bother the original author with issues from this fork. 

[![Discord](https://img.shields.io/badge/Join-Discord-blue?logo=discord&logoSize=auto)](https://discord.gg/222yVp6pUq)


---
## 🚀 Features Implemented
```
- Login  
- Newbie tutorial  
- Team management  
- Gacha  
- Story (pending testing)  
- Basic account management  
- MomoTalk  
- Mail global/personal management  
- Character growth management  
- Inventory management  
- Dungeons - Bounty / Commission / Scrimmage / Joint Firing Drill  
- Auto restoration of recoverable items  
- Cafe  
- Friend management  
- Lesson  
- Club  
- Battle Assistance  
- Total Assault  
- Daily login rewards  
- Final Restriction  
- Grand Assault
- Shop
```
---
## 🛠️ Usage

#### The user is encouraged to compile from source. Check out build workflows for an idea on what to do.
#### For resources to generate Excel.bin, run fetch_resources.sh (only Linux and WSL2 is supported at this time)

#### Friendly hint: You can fork this fork to run the build workflows in this fork (Build.yml)

---

## ⚙️ Configuration Instructions
> Note that comments cannot exist in the actual json file
```
{
  "LogLevel": "info",
  "ResourcesPath": "./resources", // Not used in release version
  "DataPath": "./data",
  "GucooingApiKey": "123456", // Key to authenticate when using API
  "AutoRegistration": true, // Auto registration
  "Tutorial": false, // Enable tutorial - incomplete
  "HttpNet": {
    "InnerAddr": "0.0.0.0", // Listening address
    "InnerPort": "5000", // Listening port
    "OuterAddr": "10.0.0.3", // External address
    "OuterPort": "5000", // External port
    "Tls": false, // Enable SSL
    "CertFile": "./data/cert.pem",
    "KeyFile":   "./data/key.pem"
  },
  "GateWay": {
    "MaxPlayerNum": 0, // Max online players
    "MaxCachePlayerTime": 720, // Max player cache time
    "BlackCmd": {}, // Not used in release version
    "IsLogMsgPlayer": true // Not used in release version
  },
  "DB": {
    "dbType": "sqlite", // Database type, supports sqlite and mysql
    "dsn": "BaPs.db" // Database address, if mysql please fill mysql url
  },
  "RaidRankDB": {
    "dbType": "sqlite", // Database type, supports sqlite and mysql
    "dsn": "Rank.db" // Database address, if mysql please fill mysql url
  },
  "Irc": { // Can use general IRC server address
    "HostAddress": "127.0.0.1", // Club chat server IRC address
    "Port": 16666, // Club chat server IRC port
    "Password": "mx123" // Club chat server IRC password
  }
}
```
---

## 🌐 Proxy Settings
Proxy the following addresses: where `http://127.0.0.1:5000` is the server address
```plaintext
https://ba-jp-sdk.bluearchive.jp  →  http://127.0.0.1:5000
https://yostar-serverinfo.bluearchiveyostar.com → http://127.0.0.1:5000
```

### ⛓️Proxy Solution

You can view the following docs
- [Android_MitmProxy Proxy Solution](Android_Mitmproxy_Readme_EN.md)

---

## ⌨️ GM Tool - [BlueArchiveGM](https://github.com/AzureXuanVerse/BlueArchiveGM)
Download from [GitHub Actions](https://github.com/AzureXuanVerse/BlueArchiveGM/actions) 

Default connection address: `http://127.0.0.1:5000` 
Default key: Random secure hash (can be changed in config.json)

---
## 🤝 Contribute
We welcome everyone who wants to help us, you can help us in the following ways:
- 🐛 Submit an Issue to report problems
- 💡 Submit a Pull Request to improve the code
- 📖 Improve project documentation
- 🚀 Join the Discord channel to provide suggestions
---

## ⚠️ Notes
1. Due to copyright reasons, the resources used in dev will not be made public
2. Due to copyright reasons, some source codes will not be made public, but we can guarantee that the non-public parts of the code have no malicious content
3. Player data will not be saved to the database in real-time, if you need the latest data, you can access player data through the API

---
## 🤜 Acknowledgements

- Thanks to [zset](https://github.com/liyiheng/zset) for implementing the leaderboard

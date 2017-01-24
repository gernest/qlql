var electron = require('electron');
// Module to control application life.
var app = electron.app;
// Module to create native browser window.
var BrowserWindow = electron.BrowserWindow;
var spawn= require('child_process').spawn;

var fs=require('fs');
var path = require('path');
var url = require('url');

// Keep a global refe   rence of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
var  mainWindow=null;

function createWindow() {
    // Create the browser window.
    mainWindow = new BrowserWindow({ width: 800, height: 600 });

    // and load the index.html of the app.
    mainWindow.loadURL(url.format({
        pathname: path.join(__dirname, 'index.html'),
        protocol: 'file:',
        slashes: true
    }));



    // Emitted when the window is closed.
    mainWindow.on('closed', function () {
        // Dereference the window object, usually you would store windows
        // in an array if your app supports multi windows, this is the time
        // when you should delete the corresponding element.
        mainWindow = null;
    });
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', createWindow);

// Quit when all windows are closed.
app.on('window-all-closed', function () {
    // On OS X it is common for applications and their menu bar
    // to stay active until the user quits explicitly with Cmd + Q
    if (process.platform !== 'darwin') {
        app.quit();
    }
});

app.on('activate', function () {
    // On OS X it's common to re-create a window in the app when the
    // dock icon is clicked and there are no other windows open.
    if (mainWindow === null) {
        createWindow();
    }
});
fs.readdir(__dirname,function(error,files){
    if (error){
        throw error;
    }
    files.forEach(function(file){
        base=path.basename(file);
        re=/ql-api-server*/;
        if(base.match(re)){
            full=path.join(__dirname,file);
            console.log(full);
            ql=spawn(full,["server","--port",8000]);
            ql.stdout.on('data',function(data){
                console.log(data.toString())
            });
            ql.stderr.on('data',function(data){
                console.log(data.toString())
            });
            ql.on('close',function(data){
                console.log(data.toString())
            });
        }
    });
});



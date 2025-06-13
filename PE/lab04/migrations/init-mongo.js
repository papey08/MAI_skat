db = db.getSiblingDB('storage_db');
db.createCollection('storages');
db.createCollection('files');
db.files.createIndex({ name: 1, user_id: 1, path: 1 });

if (!db.storages.findOne({ user_id: 1 })) {
    db.storages.insertOne({
        user_id: 1,
        root: { folders: {}, files: {} }
    });
}

let storage = db.storages.findOne({ user_id: 1 });
let root = storage.root;

function createFolder(path) {
    let pathParts = path.split('/');
    let currentFolder = root;

    for (let part of pathParts) {
        if (!currentFolder.folders[part]) {
            currentFolder.folders[part] = { folders: {}, files: {} };
        }
        currentFolder = currentFolder.folders[part];
    }
}

createFolder('aaa');
createFolder('aaa/bbb');
createFolder('aaa/ccc');

db.storages.updateOne({ user_id: 1 }, { $set: { root: root } });

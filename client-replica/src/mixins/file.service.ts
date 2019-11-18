import { Component, Mixins } from 'vue-property-decorator';

import { DataService } from './data.service';

import { FileInfo } from '@/interfaces/file.interface';

interface FileList {
  files: FileInfo[]
}

@Component
export class FileService extends Mixins(DataService) {

  async getFiles() {
    let obs = await this.getData("/api/files");
    return new Promise<FileList>((resolve, reject) => {
      obs.subscribe(
        (fileList: FileList) => {
          resolve(fileList);
        },
        (error) => {
          reject(error);
        }
      )
    });
  }

}
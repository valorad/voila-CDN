import { Component, Mixins } from 'vue-property-decorator';

import { DataService } from './data.service';

import { HostInfo } from '@/interfaces/host.interface';

@Component
export class HostService extends Mixins(DataService) {

  async getHostInfo() {

    let obs = await this.getData("/api/ip");
    return new Promise<HostInfo>((resolve, reject) => {
      obs.subscribe(
        (hostInfo) => { resolve(hostInfo); },
        (error) => {reject(error); }
      );
    });
  }

}
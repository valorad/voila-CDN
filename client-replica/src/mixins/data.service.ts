import { Component, Vue } from 'vue-property-decorator';
import { distinctUntilChanged, catchError } from 'rxjs/operators';
import { of } from 'rxjs';

import Axios from 'axios';

@Component
export class DataService extends Vue {
  
  private request = async (url: string, extractMethod?: (value: any, index: number) => any) => {


    try {

      let response = await Axios.get(url);

      if (extractMethod) {
        this.extractData(response.data, extractMethod)
      }

      return response;

    } catch (error) {
      console.error(error);
      return {} as any;
    }

  };

  private observe = (obj: any) => {

    return of(obj)
            .pipe(
              distinctUntilChanged(),
              catchError((error) => of(error || "Server Error"))
            )

  };

  private extractData = (data: any, extractMethod: (value: any, index: number) => any) => {

      if (data instanceof Object && !(data instanceof Array)) {
        data = [data];
      }

      data.map(extractMethod);

      if (data.length <= 1) {
        data = data[0];
      }

  };


  getData = async (url: string, extractMethod?: (value: any, index: number) => any) => {

    let response = await this.request(url, extractMethod);
    let data = response.data;
    
    return this.observe(data);

  };

  getResponse = async (url: string, extractMethod?: (value: any, index: number) => any) => {

    let response = await this.request(url, extractMethod);

    return this.observe(response);

  };

  postData = async (url: string, data: any) => {

    let response = await Axios.post(url, data);

    return this.observe(response.data);

  };

  patchData = async (url: string, data: any) => {

    let response = await Axios.patch(url, data);

    return this.observe(response.data);

  };

  deleteData = async (url: string) => {

    let response = await Axios.delete(url);

    return this.observe(response.data);

  };

}
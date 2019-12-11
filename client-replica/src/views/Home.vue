<template>
  <section class="home">
    <main class="home">
    <div class="contentBox">
      <md-tabs class="tabs">
        <md-tab id="tab-files" md-label="Files">
          
            <section class="allFiles">
              <md-table md-card>
                <md-table-toolbar>
                  <h1 class="md-title">Files</h1>
                </md-table-toolbar>

                <md-table-row class="head">
                  <md-table-head md-numeric>ID</md-table-head>
                  <md-table-head>Name</md-table-head>
                  <md-table-head>Type</md-table-head>
                  <md-table-head>Actions</md-table-head>
                </md-table-row>

                <md-table-row v-for="(file, index) of files" :key="`${index}-${file.name}`"> 
                  <md-table-cell md-numeric>{{ index + 1 }}</md-table-cell>
                  <md-table-cell>
                    <md-icon>insert_drive_file</md-icon>
                    <a :href="`${remoteMediaPath}/${encodeName(file.name)}`">{{ file.name }}</a>
                  </md-table-cell>
                  <md-table-cell>{{ file.type }}</md-table-cell>
                  <md-table-cell> <a :href="`${remoteMediaPath}/${encodeName(file.name)}`"><md-icon>cloud_download</md-icon></a> </md-table-cell>
                </md-table-row>
              </md-table>

              <h1 class="md-title" v-if="files.length <= 0">No files found on this server.</h1>
            </section>
          
        </md-tab>
        <md-tab id="tab-videos" md-label="Videos">

          <section class="videos">
            <ul>

              <li class="video" v-if="videos.length <= 0">
                <div class="player">
                  <h2>No videos available on this server.</h2>
                </div>
              </li>

              <li v-for="(video, index) of videos" :key="`${index}-${video.name}`" class="video">
                <div class="player">
                  <h2>{{ video.fileName }}</h2>
                  <video-player
                    class="vjs-custom-skin"
                    :options="video.options"
                    :playsinline="true"
                  >
                  </video-player>
                </div>
              </li>

            </ul>
          </section>



          <div class="player">

          </div>
        </md-tab>
      </md-tabs>
    </div>
    </main>
  </section>
</template>

<style lang="scss" scoped>
@import "src/var";

.contentBox {
  max-width: 1600px;
  margin: 0 auto;
  @extend %fCentered;
}

.tabs {
  flex-grow: 1;
}

section.allFiles {
  flex-grow: 1;
}

section.videos {
  max-width: 108em;
  margin: 1em auto;
  ul {
    display: flex;
    flex-wrap: wrap;
  }
  li {
    margin: 1em auto;
    width: 47%;
  }
  .player {
    max-height: 30em;
  }
}



</style>

<script lang="ts">
import { Component, Mixins } from 'vue-property-decorator';
import { videoPlayer } from 'vue-video-player';

import { FileService } from "../mixins/file.service";

import { FileInfo } from '../interfaces/file.interface';

interface VideoFile {
  fileName: string,
  options: any,
}

@Component({
  components: {
    videoPlayer
  }
})
export default class Home extends Mixins(FileService) {

  remoteMediaPath = "/media";
  // remoteMediaPath = "/statics";

  files: FileInfo[] = [];

  videos: VideoFile[] = [];

  supportedVideoTypes = ["video/mp4", "video/mpeg4", "video/webm", "video/ogg"]

  playerBaseOptions = {
    autoplay: false,
    muted: false,
    language: 'en',
    playbackRates: [0.7, 1.0, 1.5, 2.0],
    height: "400px",
  }

  async init() {
    this.files = await this.extractFileList();
    this.filterVideoFiles();
  }

  encodeName(name: string) {
    return encodeURIComponent(name);
  }

  async extractFileList() {

    let fileList = await this.getFiles();

    return fileList.files;
  }

  filterVideoFiles() {
    let filteredVideos = this.files.filter((file) => {
      return this.supportedVideoTypes.includes(file.type);
    })

    for (let video of filteredVideos) {
      this.videos.push({
        fileName: video.name,
        options: {
          ...this.playerBaseOptions,
          sources: [{src: `${this.remoteMediaPath}/${this.encodeName(video.name)}`}]
        }
      });
    }


  }


  mounted() {
    this.init()
  }

}
</script>


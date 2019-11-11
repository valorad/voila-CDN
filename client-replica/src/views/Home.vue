<template>
  <section class="home">
    <main class="home">
      <div class="contentBox">
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
        </section>

      </div>

    </main>
  </section>
</template>

<style lang="scss" scoped>
@import "src/var";

.contentBox {
  max-width: 1400px;
  margin: 0 auto;
  @extend %fCentered;
}

section.allFiles {
  flex-grow: 1;
}

</style>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

interface FileInfo {
  name: string,
  type: string
}

@Component
export default class Home extends Vue {

  remoteMediaPath = "/media";

  files: FileInfo[] = []

  async init() {
    let res =         {
            "files": [
                {
                    "name": ".gitkeep",
                    "type": "inode/x-empty"
                },
                {
                    "name": "amd.txt",
                    "type": "text/plain"
                },
                {
                    "name": "TV6.rmvb",
                    "type": "application/vnd.rn-realmedia-vbr"
                },
                {
                    "name": "Cooking Class.7z",
                    "type": "application/x-7z-compressed"
                },
                {
                    "name": "Phonetiques.flv",
                    "type": "application/x-shockwave-flash"
                },
                {
                    "name": "AppleWWDC2019.mp4",
                    "type": "video/mpeg4"
                }
            ]
        }
    
    this.files = res.files;

  }

  encodeName(name: string) {
    return encodeURIComponent(name);
  }

  mounted() {
    this.init()
  }

}
</script>


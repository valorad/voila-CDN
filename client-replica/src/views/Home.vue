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

            <md-table-row> 
              <md-table-cell md-numeric>1</md-table-cell>
              <md-table-cell> <md-icon>insert_drive_file</md-icon> <a href="##">.gitkeep</a> </md-table-cell>
              <md-table-cell>inode/x-empty</md-table-cell>
              <md-table-cell> <a href="##"><md-icon>cloud_download</md-icon></a> </md-table-cell>
            </md-table-row>

            <md-table-row>
              <md-table-cell md-numeric>2</md-table-cell>
              <md-table-cell><md-icon>insert_drive_file</md-icon><a href="##">Cooking Class.7z</a></md-table-cell>
              <md-table-cell>application/x-7z-compressed</md-table-cell>
              <md-table-cell><a href="##"><md-icon>cloud_download</md-icon></a></md-table-cell>
            </md-table-row>

            <md-table-row>
              <md-table-cell md-numeric>3</md-table-cell>
              <md-table-cell><md-icon>insert_drive_file</md-icon><a href="##">AppleWWDC2019.mp4</a></md-table-cell>
              <md-table-cell>video/mpeg4</md-table-cell>
              <md-table-cell><a href="##"><md-icon>cloud_download</md-icon></a></md-table-cell>
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

@Component
export default class Home extends Vue {}
</script>


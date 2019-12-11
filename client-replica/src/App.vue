<template>
  <section id="app">
    <md-app class="appHolder">
      <md-app-toolbar class="md-primary navbar">

        <div class="menuHolder">
          <md-button class="md-icon-button" @click="toggleNav()">
            <md-icon>menu</md-icon>
          </md-button>
        </div>

        <header>
          <router-link :to="`/index`">
            <span class="md-title">Voila-CDN-Replica</span>
          </router-link>
        </header>
        
        <!-- <nav>
          
            <md-button class="md-dense">
              index
            </md-button>
          

        </nav> -->
        <div class="flexSpacer"></div>
        <footer>
          <div class="ip">
            <ul v-if="hostInfo">
              <li> <md-icon>computer</md-icon> <span v-html="hostInfo.host"></span> </li>
              <li> <md-icon>my_location</md-icon> <span v-html="hostInfo.ip"></span> </li>
            </ul>
          </div>
          <div class="github">
            <a target="_blank" rel="noopener noreferrer" href="https://github.com/valorad/voila-CDN"><i class="iconfont icon-github"></i></a>
          </div>

        </footer>
      </md-app-toolbar>

      <md-app-drawer :md-active.sync="isNavShown" md-swipeable>
        <md-toolbar class="md-transparent" md-elevation="0">
          Navigation
        </md-toolbar>
        <md-list>
          <md-list-item>
            <md-icon>move_to_inbox</md-icon>
            <span class="md-list-item-text">Inbox</span>
          </md-list-item>
        </md-list>
      </md-app-drawer>

      <md-app-content>

        <section class="router">
          <router-view />
        </section>
        
      </md-app-content>

    </md-app>
  </section>
</template>

<style lang="scss">

@import url("https://fonts.googleapis.com/css?family=Roboto:400,500,700,400italic|Material+Icons");
@import url("//at.alicdn.com/t/font_1492360_rcvfc2l2j9h.css");

// vue material
@import "~vue-material/dist/vue-material.min.css";
@import "themeApplied";

// vue video player
@import "~video.js/dist/video-js.min.css";
@import "~vue-video-player/src/custom-theme.css";

// custom
@import "var";

#app {
  font-family: 'Roboto', '微软雅黑', Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.flexSpacer {
  flex-grow: 1;
}

.appHolder {
  width: 100vw;
  height: 100vh;
}

.navbar {
  display: flex;

  i {
    font-size: 3em;
    @extend %fCentered;
    color: $foreground;
  }

  header {
    @extend %fCentered;
  }

  nav {
    display: none;
    a {
      display: block;
      height: 100%;
      width: 100%
    }
  }

  footer {
    display: flex;
  }

  .ip {
    @extend %fCentered;
    ul {
      @extend %fCentered;
    }
    li {
      @extend %fCentered;
      margin-right: 1.5em;
    }
    i {
      margin: 0 0.5em;
    }
  }

  .gitlab {
    @extend %fCentered;
    margin-left: 1em;
  }

}

.router {
  height: 100%;
}

ul {
  list-style-type: none;
  margin: 0;
  padding: 0;
}

@media only screen and (min-width: 768px) {
  .navbar {
    .menuHolder {
      display: none;
    }
    nav {
      display: flex;
    }
  }
  
}

</style>

<script lang="ts">
import { Component, Mixins } from 'vue-property-decorator';
import { HostInfo } from "./interfaces/host.interface";

import { HostService } from "./mixins/host.service";

@Component
export default class App extends Mixins(HostService) {
  isNavShown = false;

  hostInfo = { } as HostInfo;

  toggleNav() {
    this.isNavShown = !this.isNavShown;
  };

  async init() {
    this.hostInfo = await this.getHostInfo();
  }

  mounted() {
    this.init();
  }

}
</script>

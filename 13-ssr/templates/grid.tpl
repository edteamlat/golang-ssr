{{ define "grid" }}

  {{ template "navigation" }}
  <div class="ed-grid m-grid-2 lg-grid-4 row-gap">

  {{ range $i, $course := .Courses }}
    <article
      class="course-card s-radius-1 s-shadow-bottom background s-shadow-card-micro s-transition white-color nowrap s-column s-mb-0">
      <div class="s-column flex-grow ">
        <div class="img-container s-ratio-16-9 s-radius-tl-1 s-radius-tr-1">
          <a href="/cursos/{{ $course.Slug }}">
            <img class="img-course-card s-transition" alt="{{ $course.Title }}"
              src="{{ $course.Img }}">
          </a>
        </div>
        <div title="{{ $course.Title }}" class="course-progress s-mt-05"></div>
        <div class="s-pxy-2 s-column flex-grow">
          <a href="/cursos/{{ $course.Slug }}">
            <h3 class="s-mb-1 t4">
              <span class="s-color-grey-700">{{ $course.Name }}</span>
            </h3>
          </a>
          <p class="description small s-row-text-3 s-mb-2"
            title="{{ $course.Description }}">
                {{ $course.Description }}
          </p>
          <div class="flex nowrap s-to-bottom">
            <a
              class="s-cross-center smaller background blue-100 s-radius s-px-1 s-py-05 s-mr-1" href="/cursos">
              <svg
                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 16.103" class="svg-icon s-mr-05 fill blue-500">
                <path
                  d="M19.945,53.621,18.1,50.152a.468.468,0,0,0-.29-.232h0l-7.671-2.071a.469.469,0,0,0-.244,0l-7.71,2.071h0a.468.468,0,0,0-.29.232L.055,53.621a.468.468,0,0,0,.214.643l1.6.757V59.76a.468.468,0,0,0,.265.422l7.681,3.707a.465.465,0,0,0,.4,0h0L17.9,60.182a.468.468,0,0,0,.265-.422V55.022c0-.006,0-.012,0-.018l1.564-.74a.468.468,0,0,0,.214-.643Zm-9.926-4.835,6.134,1.656L10,52.722,3.85,50.442ZM2.531,50.953,9.37,53.487,7.864,56.818,1.112,53.627Zm.276,8.513v-4l5.087,2.4a.468.468,0,0,0,.627-.23l1.03-2.277v7.359Zm14.424,0-6.744,3.255V55.447l.991,2.192a.468.468,0,0,0,.627.23l5.126-2.423Zm-5.1-2.648L10.63,53.487l6.839-2.535,1.419,2.674Z"
                  transform="translate(0 -47.832)"></path>
              </svg>
              <span class="s-grey-text-color">Curso</span>
            </a>
            <div class="s-cross-center small ">
              <div class="flex nowrap s-cross-center s-mr-1">
                <svg xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 49.94 49.94" class="svg-icon s-mr-05 fill accent-color ">
                  <path
                    d="M48.856,22.73c0.983-0.958,1.33-2.364,0.906-3.671c-0.425-1.307-1.532-2.24-2.892-2.438l-12.092-1.757c-0.515-0.075-0.96-0.398-1.19-0.865L28.182,3.043c-0.607-1.231-1.839-1.996-3.212-1.996c-1.372,0-2.604,0.765-3.211,1.996L16.352,14c-0.23,0.467-0.676,0.79-1.191,0.865L3.069,16.622c-1.359,0.197-2.467,1.131-2.892,2.438c-0.424,1.307-0.077,2.713,0.906,3.671l8.749,8.528c0.373,0.364,0.544,0.888,0.456,1.4L8.224,44.701c-0.183,1.06,0.095,2.091,0.781,2.904c1.066,1.267,2.927,1.653,4.415,0.871l10.814-5.686c0.452-0.237,1.021-0.235,1.472,0l10.815,5.686c0.526,0.277,1.087,0.417,1.666,0.417c1.057,0,2.059-0.47,2.748-1.288c0.687-0.813,0.964-1.846,0.781-2.904l-2.065-12.042c-0.088-0.513,0.083-1.036,0.456-1.4L48.856,22.73z">
                  </path>
                </svg>
                <p class="s-mb-0 s-color-text-light">{{ $course.Average }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <footer
        class="background grey-color-lighter s-px-2 s-py-1 s-radius-bl-1 s-radius-br-1 flex nowrap s-z-normal s-to-bottom footer-card s-relative">
        <div class="s-main-center ">
          <div title="{{ $course.Professor }}" class="card__teacher s-cross-center nowrap ">
            <div id="" class="user-avatar flex-none s-small s-mr-1 ">
              <div class="img-container circle">
                <img
                  src="{{ $course.ProfessorImg }}"
                  alt="Avatar" loading="lazy">
              </div>
            </div>
              <span class="card__teacher-name smaller s-row-text-1 s-z-normal s-mr-1">
                {{ $course.Professor }}
              </span>
          </div>
        </div>
        <div class="s-to-right">
          <div role="presentation" class="s-cross-center s-cursor-pointer">
            <div class="">
              <button
                class="small s-border-none background transparent-color s-cursor-pointer s-cross-center s-nowrap">
                  <svg
                    xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32" class="svg-icon normal s-mr-05 fill blue-500">
                    <g>
                    <path
                      d="M24.0003 26.6177C23.9997 27.1872 24.1966 27.7394 24.5575 28.18C24.9183 28.6207 25.4209 28.9225 25.9794 29.0342C26.5379 29.1458 27.1178 29.0604 27.6204 28.7923C28.123 28.5243 28.5171 28.0903 28.7355 27.5642C28.9539 27.0382 28.9832 26.4527 28.8183 25.9076C28.6534 25.3624 28.3046 24.8913 27.8313 24.5745C27.3579 24.2577 26.7894 24.1148 26.2225 24.1702C25.6556 24.2256 25.1255 24.4758 24.7225 24.8783C24.4897 25.1035 24.3056 25.374 24.1813 25.6732C24.0571 25.9723 23.9955 26.2938 24.0003 26.6177Z">
                    </path>
                    <rect x="7" y="2" width="2" height="7" rx="1" transform="rotate(90 7 2)"></rect>
                    <rect x="11.0105" y="21.7759" width="2" height="20.1162" rx="1"
                      transform="rotate(167.811 11.0105 21.7759)"></rect>
                    <rect x="8.93015" y="24.5331" width="2" height="4.07966" rx="1"
                      transform="rotate(-150 8.93015 24.5331)"></rect>
                    <rect x="8" y="24" width="2" height="21" rx="1" transform="rotate(-90 8 24)"></rect>
                    <rect x="32" y="4" width="2" height="26" rx="1" transform="rotate(90 32 4)"></rect>
                    <rect x="32" y="17" width="2" height="13" rx="1" transform="rotate(-180 32 17)"></rect>
                    <rect x="31.84" y="15" width="2" height="23.8096" rx="1" transform="rotate(83.3797 31.84 15)">
                    </rect>
                    <path
                      d="M7.3851 26.6173C7.38504 27.187 7.5825 27.739 7.94385 28.1793C8.3052 28.6196 8.80806 28.921 9.36673 29.0321C9.9254 29.1431 10.5053 29.0571 11.0076 28.7885C11.5099 28.5199 11.9035 28.0854 12.1214 27.5591C12.3392 27.0328 12.3679 26.4472 12.2023 25.9022C12.0368 25.3572 11.6874 24.8864 11.2136 24.5702C10.7399 24.2539 10.1711 24.1117 9.6043 24.1678C9.03746 24.2239 8.50763 24.4749 8.10509 24.8779C7.87264 25.1033 7.68883 25.3739 7.56499 25.6731C7.44116 25.9722 7.37994 26.2936 7.3851 26.6173V26.6173Z">
                    </path>
                    </g>
                  </svg>
                  <span class="color s-font-semibold blue-500 s-whitespace-nowrap">
                    {{ usd $course.Price}}
                  </span>
                </button>
              </div>
          </div>
        </div>
      </footer>
    </article>
  
    {{ end }}
 
  </div>

{{ end }}
const slider = document.getElementById("volumeSlider");
const volumeValueText = document.getElementById("volumeValue");

slider.oninput = function() {
    volumeValueText.innerHTML = this.value + "%";
}

function startFfmpeg() {
    const btn = document.querySelector('.btn-start');
    btn.innerHTML = '<i class="fas fa-spinner fa-spin"></i> Starting...';
    btn.style.opacity = '0.8';
    btn.style.cursor = 'not-allowed';
    setTimeout(() => {
        btn.innerHTML = '<i class="fas fa-stop"></i> Stop FFmpeg Listen';
        btn.style.background = 'linear-gradient(135deg, #ef4444, #dc2626)';
        btn.style.opacity = '1';
        btn.style.cursor = 'pointer';
    }, 2000);
}
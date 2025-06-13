using UnityEngine;

[RequireComponent(typeof(Renderer))]
public class EnemyAppearanceAudio : MonoBehaviour
{
    public AudioClip appearanceClip;

    private static AudioSource globalAudioSource;
    private static EnemyAppearanceAudio currentlyPlaying;

    private Renderer enemyRenderer;
    private bool isVisible = false;
    private bool wasPaused = false;

    void Start()
    {
        enemyRenderer = GetComponent<Renderer>();

        if (globalAudioSource == null)
        {
            GameObject audioObj = new GameObject("EnemyAudioSource");
            globalAudioSource = audioObj.AddComponent<AudioSource>();
            globalAudioSource.loop = true;
            globalAudioSource.playOnAwake = false;
            DontDestroyOnLoad(audioObj);
        }
    }

    void Update()
    {
        if (Time.timeScale == 0)
        {
            if (globalAudioSource.isPlaying)
            {
                globalAudioSource.Pause();
                wasPaused = true;
            }
        }
        else if (wasPaused)
        {
            if (currentlyPlaying == this)
            {
                globalAudioSource.UnPause();
            }
            wasPaused = false;
        }

        bool nowVisible = enemyRenderer.isVisible;

        if (nowVisible && !isVisible)
        {
            isVisible = true;
            PlayAppearanceSound();
        }
        else if (!nowVisible && isVisible)
        {
            isVisible = false;
            StopIfThisEnemy();
        }
    }

    void OnDestroy()
    {
        StopIfThisEnemy();
    }

    void PlayAppearanceSound()
    {
        if (currentlyPlaying != null && currentlyPlaying != this)
        {
            currentlyPlaying.StopIfThisEnemy();
        }

        if (appearanceClip != null)
        {
            globalAudioSource.clip = appearanceClip;
            globalAudioSource.Play();
            currentlyPlaying = this;
        }
    }

    void StopIfThisEnemy()
    {
        if (currentlyPlaying == this)
        {
            globalAudioSource.Stop();
            currentlyPlaying = null;
        }
    }
}

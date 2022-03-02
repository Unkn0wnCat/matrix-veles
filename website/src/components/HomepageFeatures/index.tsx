import React from 'react';
import clsx from 'clsx';
import styles from './styles.module.css';

type FeatureItem = {
  title: string;
  image: string;
  description: JSX.Element;
};

const FeatureList: FeatureItem[] = [
  {
    title: 'Lightweight',
    image: null,
    description: (
      <>
        Veles is built to be light on storage, memory and CPU. Run it
        on your server, PC, Raspberry Pi or Smart Toaster!
      </>
    ),
  },
  {
    title: 'Modern Codebase',
    image: null,
    description: (
      <>
        Veles is built from the ground up to deliver you the best
        experience using GoLang - a next-gen language from Google.<br/>
        And the best thing? It's open source! Fork it, make a mod where
        all bad messages are replaced by cat images! You are free!
      </>
    ),
  },
  {
    title: 'Convenient Web Interface',
    image: null,
    description: (
      <>
        Veles can be managed from any PC in the world. All you
        need to access the modern web interface is an internet
        connection and a web browser.
      </>
    ),
  },
];

function Feature({title, image, description}: FeatureItem) {
  return (
    <div className={clsx('col col--4')}>
        {/*<div className="text--center">
        <img className={styles.featureSvg} alt={title} src={image} />
      </div>*/}
      <div className="text--center padding-horiz--md">
        <h3>{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): JSX.Element {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}

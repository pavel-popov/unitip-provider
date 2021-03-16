#!/usr/local/bin/bb
(ns unitip
  (:require [cheshire.core :as json]))


(defn aws-s3-link [input]
  (when-let [[_ bucket key] (re-matches #"s3://([\w\-_\d\.]+)/([\w\-_\d\.\/]+.html)" input)]
    {:type  :url
     :label "S3 URL"
     :autoExecuteIfFirst true
     :value (str "https://" bucket ".s3.amazonaws.com/" key)}))


(def aws-batch-job-url-prefix "https://eu-west-1.console.aws.amazon.com/batch/v2/home?region=eu-west-1#/jobs/detail/")


(defn aws-batch-link [input]
  (when-let [[_ job-id] (re-matches #".*\"JobId\": \"([a-z0-9\-]{36})\".*" input)]
    {:type  :url
     :label (str "AWS Batch GUIv2 for job " job-id)
     :autoExecuteIfFirst true
     :value (str aws-batch-job-url-prefix job-id)}))


(defn aws-batch-link-by-uuid [input]
  (when-let [[_ job-id] (re-matches #"^([0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12})$" input)]
    {:type  :url
     :label (str "AWS Batch GUIv2 for job " job-id)
     :value (str aws-batch-job-url-prefix job-id)}))


(defn aws-sfn-link [input]
  (when-let [[_ region account execution]
             (re-matches #".*arn:aws:states:(.*):([0-9]*):execution:(.*).*" input)]
    {:type  :url
     :label (str "SFN GUI for execution " execution)
     :autoExecuteIfFirst true
     :value (str "https://" region ".console.aws.amazon.com/states/home?region=" region
                 "#/executions/details/arn:aws:states:" region ":"
                 account ":execution:" execution)}))


(defn google [input]
  {:type  :url
   :label "Open Google"
   :value (str "https://google.com/search?q=" input)})


(defn clojuredocs [input]
  (when-let [[_ ns func]
             (re-matches #"(.+/)?(.+)" input)]
    (let [cns (or ns "clojure.core/")]
      {:type  :url
       :label "ClojureDocs"
       :value (str "https://clojuredocs.org/" cns func)})))


(import 'java.time.format.DateTimeFormatter
        'java.time.LocalDateTime
        'java.time.Instant)


(defn unix-timestamp [input]
  (when-let [ts (Instant/ofEpochSecond (Long/valueOf input))]
    {:type :text
     :value (.toString ts)}))


;; (unix-timestamp 1615588800)


;; (clojuredocs "clojure.str/join")

;; (clojuredocs "asdf")

;; (google "clojure.core/asdf")


(defn handle [input]
  (map
   (fn [handler]
     (try
       (handler input)
       (catch Exception e
         nil)))
   [aws-s3-link
    aws-batch-link
    aws-batch-link-by-uuid
    aws-sfn-link
    google
    clojuredocs
    unix-timestamp]))


(->> *command-line-args*
    first
    handle
    (filter some?)
    json/encode
    println)


;; (re-matches #"s3://([\w\-_\d\.]+)/([\w\-_\d\.\/]+)" "s3://fh-search/metrics/lines-to-sqs/s3--fh-search-prod-tags-hotel-2020-12-21-00.json.gz.html")
